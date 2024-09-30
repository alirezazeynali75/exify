package outbox

import (
	"context"
	"log/slog"
	"time"

	"github.com/alirezazeynali75/exify/internal/eventbus"
)

type eventBus interface {
	Produce(ctx context.Context, msgs []eventbus.MessageToPublish) error
}

type outboxRepo interface {
	GetPendingEventsAndUpdateStatus(ctx context.Context) ([]Outbox, error)
	UpdateByID(ctx context.Context, id uint64, status OutboxStatus) error
	BatchUpdatePendingBasedOnTime(ctx context.Context, status OutboxStatus, maxAge time.Duration) error
}

type OutboxService struct {
	logger *slog.Logger
	eventBus eventBus
	outboxRepo outboxRepo
}

func NewOutboxService(
	logger *slog.Logger,
	eventBus eventBus,
	outboxRepo outboxRepo,
) *OutboxService {
	return &OutboxService{
		logger: logger,
		eventBus: eventBus,
		outboxRepo: outboxRepo,
	}
}

func (svc OutboxService) ProduceMessages(ctx context.Context) error {
	logger := svc.logger.With(slog.String("op", "ProduceMessages"))
	pendingRecord, err := svc.outboxRepo.GetPendingEventsAndUpdateStatus(ctx)
	if err != nil {
		return err
	}
	msgToPublish := make([]eventbus.MessageToPublish, len(pendingRecord))
	for i, r := range pendingRecord {
		msgToPublish[i] = eventbus.MessageToPublish{
			Topic: r.Topic,
			Value: r.Payload,
		}
	}
	err = svc.eventBus.Produce(ctx, msgToPublish)
	if err != nil {
		return err
	}
	for _, r := range pendingRecord {
		err := svc.outboxRepo.UpdateByID(ctx, r.ID, SENT)
		if err != nil {
			logger.With(slog.String("err", err.Error())).Error("there is an error")
		}
	}
	return nil
}

func (svc OutboxService) RevertPending(ctx context.Context) error {
	return svc.outboxRepo.BatchUpdatePendingBasedOnTime(ctx, READY, time.Minute * 20)
}