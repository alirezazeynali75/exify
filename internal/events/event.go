package events

import "time"

type NewDepositEvent struct {
	Type       string    `json:"type"`
	ID         string    `json:"id"`
	TrackingId string    `json:"tracking_id"`
	IBAN       string    `json:"IBAN"`
	Gateway    string    `json:"gateway"`
	Amount     string    `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProcessingEvent struct {
	Type       string `json:"type"`
	TrackingId string `json:"tracking_id"`
}

type FinishedEvent struct {
	Type       string `json:"type"`
	TrackingId string `json:"tracking_id"`
	Status     string `json:"status"`
}

type Event[T any] struct {
	ID        string    `json:"id"`
	Payload   T         `json:"payload"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}
