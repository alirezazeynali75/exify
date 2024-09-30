package payment

import (
	"time"

	"github.com/alirezazeynali75/exify/internal/events"
	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/google/uuid"
)

type Deposit struct {
	ID         string
	TrackingId string
	IBAN       string
	Gateway    string
	Amount     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}


func FromNewDepositDtoToDepositModel(d dto.NewDepositDTO) Deposit {
	return Deposit{
		ID: uuid.NewString(),
		TrackingId: d.TrackingId,
		IBAN: d.IBAN,
		Gateway: d.Gateway,
		Amount: d.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (d *Deposit) GetEvent() (events.Event[events.NewDepositEvent], string) {
	return events.Event[events.NewDepositEvent]{
		ID: uuid.NewString(),
		Version: "1",
		CreatedAt: time.Now(),
		Payload: events.NewDepositEvent{
			Type: "DEPOSIT_INITIATED",
			ID: d.ID,
			TrackingId: d.TrackingId,
			IBAN: d.IBAN,
			Gateway: d.Gateway,
			Amount: d.Amount,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
	}, "events"
}