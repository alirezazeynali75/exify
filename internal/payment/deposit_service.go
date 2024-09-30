package payment

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"gorm.io/gorm"
)

type depositAgent interface {
	GetName() string
	GetStatus() bool
}

type DepositService struct {
	logger        *slog.Logger
	session       db.Session
	depositAgents []depositAgent
	inboxRepo     inboxRepo
	depositRepo   depositRepo
	outboxRepo    outboxRepo
}

func (svc *DepositService) ListAllAvailableAgents(ctx context.Context) ([]string, error) {
	agents := make([]string, 0)
	for _, agent := range svc.depositAgents {
		if agent.GetStatus() {
			agents = append(agents, agent.GetName())
		}
	}

	if len(agents) == 0 {
		return agents, ErrNoDepositAgentIsAvailable
	}

	return agents, nil
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

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil
	}
	return err
}
