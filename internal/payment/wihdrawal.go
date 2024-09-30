package payment

import (
	"time"

	"github.com/alirezazeynali75/exify/internal/events"
	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/shopspring/decimal"
	"github.com/google/uuid"
)

type Withdrawal struct {
	ID          string
	TrackingId  string
	Destination string
	Gateway     string
	Amount      decimal.Decimal
	Status      PaymentStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Withdrawal) GetProcessingEvent()  (events.Event[events.ProcessingEvent], string) {
	return events.Event[events.ProcessingEvent]{
		ID: uuid.NewString(),
		Version: "1",
		CreatedAt: time.Now(),
		Payload: events.ProcessingEvent{
			Type: "WITHDRAWAL_PROCESSING",
			TrackingId: t.TrackingId,
		},
	}, "events"
}

func (t *Withdrawal) GetFinishedEvent()  (events.Event[events.FinishedEvent], string) {
	return events.Event[events.FinishedEvent]{
		ID: uuid.NewString(),
		Version: "1",
		CreatedAt: time.Now(),
		Payload: events.FinishedEvent{
			Type: "WITHDRAWAL_FINISHED",
			TrackingId: t.TrackingId,
			Status: string(t.Status),
		},
	}, "events"
}

func FromNewWithdrawalDtoToTransaction(d dto.NewWithdrawalDto) Withdrawal {
	return Withdrawal{
		ID:          d.ID,
		Destination: d.Destination,
		Amount:      d.Amount,
	}
}

func FromUpdateWithdrawalDtoToTransaction(d dto.UpdateWithdrawalStatusDTO) Withdrawal {
	status := COMPLETED
	if !d.IsSuccess {
		status = FAILED
	}
	return Withdrawal{
		TrackingId: d.TrackingId,
		Status: status,
	}
}
