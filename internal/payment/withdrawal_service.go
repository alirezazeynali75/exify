package payment

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/alirezazeynali75/exify/pkg/sql"
)



type WithdrawalService struct {
	logger          *slog.Logger
	session         db.Session
	outboxRepo      outboxRepo
	paymentGateways []paymentGateway
	inboxRepo       inboxRepo
	transactionRepo withdrawalRepo
}

func NewWithdrawalService(
	logger *slog.Logger,
	session db.Session,
	outboxRepo outboxRepo,
	inboxRepo inboxRepo,
	transactionRepo withdrawalRepo,
	paymentGateways ...paymentGateway,
) *WithdrawalService {
	return &WithdrawalService{
		logger:          logger.With(slog.String("module", "payment")),
		session:         session,
		outboxRepo:      outboxRepo,
		inboxRepo:       inboxRepo,
		transactionRepo: transactionRepo,
		paymentGateways: paymentGateways,
	}
}

func (svc *WithdrawalService) getGateway(tx Withdrawal) paymentGateway {
	for _, g := range svc.paymentGateways {
		if g.CanDo(tx) {
			return g
		}
	}
	return nil
}

func (svc *WithdrawalService) AddNewWithdrawTransaction(ctx context.Context, d dto.NewWithdrawalDto) error {
	logger := svc.logger.With(slog.String("op", "AddNewTransaction"), slog.String("txId", d.ID))
	logger.Debug("going to handle new transaction")
	tx := FromNewWithdrawalDtoToTransaction(d)
	gateway := svc.getGateway(tx)
	if gateway == nil {
		logger.Error("there is no provider to handle the request")
		return ErrNoProviderFound
	}
	txData := tx

	txData.Status = PROCESSING

	txData.Gateway = gateway.GetName()

	err := svc.session.Transaction(ctx, func(ctx context.Context) error {
		err := svc.inboxRepo.InsertEvent(ctx, d.EventId)
		if err != nil {
			return err
		}
		return svc.transactionRepo.CreateNewTransaction(ctx, txData)
	})
	if sql.IsDuplicateEntry(err) {
		svc.logger.Warn("this event has been processed")
		return nil
	}

	if err != nil {
		return err
	}

	trackingId, err := gateway.Execute(ctx, tx)
	if err != nil {
		return err
	}

	tx.TrackingId = trackingId

	event, topic := tx.GetProcessingEvent()

	stringifyEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return svc.session.Transaction(ctx, func(ctx context.Context) error {
		err := svc.transactionRepo.UpdateTrackingId(ctx, tx.ID, trackingId)
		if err != nil {
			return err
		}
		return svc.outboxRepo.InsertNewEvent(ctx, string(stringifyEvent), topic)
	})
}

func (svc WithdrawalService) UpdateWithdrawalStatus(ctx context.Context, d dto.UpdateWithdrawalStatusDTO) error {
	logger := svc.logger.With(slog.String("op", "UpdateTransactionStatus"), slog.String("id", d.EventId))
	logger.Debug("going to update status")

	tx := FromUpdateWithdrawalDtoToTransaction(d)

	event, topic := tx.GetFinishedEvent()

	stringifyEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = svc.session.Transaction(ctx, func(ctx context.Context) error {
		err := svc.inboxRepo.InsertEvent(ctx, d.EventId)
		if err != nil {
			return err
		}
		err = svc.transactionRepo.UpdateStatusByTrackingId(ctx, d.TrackingId, tx.Status)
		if err != nil {
			return err
		}
		return svc.outboxRepo.InsertNewEvent(ctx, string(stringifyEvent), topic)
	})

	if sql.IsDuplicateEntry(err) {
		return nil
	}

	return err
}
