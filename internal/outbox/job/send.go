package job

import (
	"context"
	"log/slog"

	"github.com/robfig/cron"
)

type outboxService interface {
	ProduceMessages(ctx context.Context) error
	RevertPending(ctx context.Context) error
}

type SendMessageJob struct {
	logger *slog.Logger
	svc outboxService
}

func NewSendMessageJob(
	logger *slog.Logger,
	svc outboxService,
) *SendMessageJob {
	return &SendMessageJob{
		logger: logger,
		svc:    svc,
	}
}

func (job *SendMessageJob) Run() {
	err := job.svc.ProduceMessages(context.Background())
	if err != nil {
		job.logger.Error("there is an error")
		return
	}
	job.logger.Info("the message sent")
}

func (job *SendMessageJob) Register(c *cron.Cron) {
	c.AddJob("@every 1m", job)
}