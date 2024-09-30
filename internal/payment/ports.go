package payment

import "context"

type outboxRepo interface {
	InsertNewEvent(ctx context.Context, event string, topic string) error
}

type inboxRepo interface {
	InsertEvent(ctx context.Context, eventId string) error
}

type withdrawalRepo interface {
	CreateNewTransaction(ctx context.Context, tx Withdrawal) error
	UpdateTrackingId(ctx context.Context, txId string, trackingId string) error
	UpdateStatusByTrackingId(ctx context.Context, trackingId string, status PaymentStatus) error
}

type depositRepo interface {
	CreateNewTransaction(ctx context.Context, tx Deposit) error
}

type paymentGateway interface {
	CanDo(tx Withdrawal) bool
	GetName() string
	Execute(ctx context.Context, tx Withdrawal) (string, error)
}
