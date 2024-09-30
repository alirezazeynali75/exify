package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	xhttp "net/http"

	"github.com/alirezazeynali75/exify/internal/config"
	session "github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/eventbus"
	"github.com/alirezazeynali75/exify/internal/inbox"
	"github.com/alirezazeynali75/exify/internal/outbox"
	"github.com/alirezazeynali75/exify/internal/outbox/job"
	outboxRepository "github.com/alirezazeynali75/exify/internal/outbox/repo"
	"github.com/alirezazeynali75/exify/internal/payment"
	"github.com/alirezazeynali75/exify/internal/payment/presentation"
	"github.com/alirezazeynali75/exify/internal/payment/repo"
	"github.com/alirezazeynali75/exify/internal/provider/a"
	"github.com/alirezazeynali75/exify/internal/provider/b"
	"github.com/robfig/cron"

	// "github.com/alirezazeynali75/exify/pkg/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {

	cnf, err := config.Configure()
	if err != nil {
		log.Fatal(err)
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	logger.Info("service has been started")

	db, err := gorm.Open(mysql.Open(cnf.Mysql.DSN), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}

	db.AutoMigrate(
		inbox.InboxModel{},
		outboxRepository.OutboxModel{},
		repo.DepositModel{},
		repo.WithdrawalModel{},
	)

	s := session.GORM(db, nil)

	inboxRepo := inbox.NewInboxRepo(
		db,
	)

	outboxRepo := outboxRepository.NewOutboxRepo(
		db,
	)

	depositRepo := repo.NewDepositRepo(
		db,
	)

	withdrawalRepo := repo.NewWithdrawalRepo(
		db,
	)

	depositService := payment.NewDepositService(
		logger,
		s,
		inboxRepo,
		depositRepo,
		outboxRepo,
	)

	// aProviderHttpClient := http.NewClient(cnf.A.ClientURL, map[string]string{}, nil)
	// aProviderHttpClientWithCB := http.NewHttpCircuitBreaker(
	// 	cnf.A.GetCircuitBreakerConfig(),
	// 	aProviderHttpClient,
	// )
	aProviderClient := a.NewAProviderMock()
	bProviderClient := b.NewBProviderMock()

	withdrawalService := payment.NewWithdrawalService(
		logger,
		s,
		outboxRepo,
		inboxRepo,
		withdrawalRepo,
		aProviderClient,
		bProviderClient,
	)

	paymentController := presentation.NewPaymentController(
		withdrawalService,
	)

	kafkaConfig, err := cnf.Kafka.ToSaramaConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	eventbus := eventbus.NewKafka(
		logger,
		cnf.Kafka.Brokers,
		kafkaConfig,
	)

	outboxService := outbox.NewOutboxService(
		logger,
		eventbus,
		outboxRepo,
	)

	sendingJob := job.NewSendMessageJob(logger, outboxService)
	revertJob := job.NewRevertStalledMessageJob(logger, outboxService)

	c := cron.New()

	sendingJob.Register(c)
	revertJob.Register(c)

	c.Start()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	gp := e.Group("/api/v1")

	paymentController.Register(gp)

	aCallbackController := a.NewACallbackController(withdrawalService, depositService)

	aCallbackController.Register(gp)

	listenAddress := fmt.Sprintf("%s:%s", cnf.Http.Address, cnf.Http.Port)

	// c := cron.New()
	go func() {
		if err := e.Start(listenAddress); err != nil && errors.Is(err, xhttp.ErrServerClosed) {
			logger.Warn("shutting down the server", slog.String("err", err.Error()))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	c.Stop()
	logger.Info("going to graceful shutdown")
	err = e.Shutdown(context.TODO())
	if err != nil {
		logger.With(slog.String("err", err.Error())).Error("couldn't stop the server")
	}
}