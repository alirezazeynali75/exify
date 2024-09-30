package payment

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/alirezazeynali75/exify/pkg/sql"
)

type DepositService struct {
	logger        *slog.Logger
	session       db.Session
	inboxRepo     inboxRepo
	depositRepo   depositRepo
	outboxRepo    outboxRepo
}

func NewDepositService(
	logger *slog.Logger,
	session db.Session,
	inboxRepo inboxRepo,
	depositRepo depositRepo,
	outboxRepo outboxRepo,
) *DepositService {
	return &DepositService{
		logger:        logger,
		session:       session,
		inboxRepo:     inboxRepo,
		outboxRepo:    outboxRepo,
		depositRepo:   depositRepo,
	}
}

func (svc *DepositService) AddDeposit(ctx context.Context, d dto.NewDepositDTO) error {
	logger := svc.logger.With(slog.String("op", "AddDeposit"), slog.String("requestId", d.RequestId))
	logger.Debug("going to add new deposit")

	deposit := FromNewDepositDtoToDepositModel(d)

	event, topic := deposit.GetEvent()

	stringifyEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = svc.session.Transaction(ctx, func(ctx context.Context) error {
		err := svc.inboxRepo.InsertEvent(ctx, d.RequestId)
		if err != nil {
			return err
		}
		err = svc.depositRepo.CreateNewTransaction(ctx, deposit)
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
