package payment

import (
	"context"
	"log/slog"
)

type paymentGateway interface {
	CanDo(tx Transaction) bool
	GetName() string
}

type PaymentService struct {
	logger *slog.Logger
	paymentGateways []paymentGateway
}

func NewPaymentService(
	logger *slog.Logger,
) *PaymentService {
	return &PaymentService{
		logger: logger.With(slog.String("module", "payment")),
	}
}

func (svc *PaymentService) getGateway(tx Transaction) paymentGateway {
	for _, g := range svc.paymentGateways {
		if g.CanDo(tx) {
			return g
		}
	}
	return nil
}

func (svc *PaymentService) AddNewTransaction(ctx context.Context, tx Transaction) error {
	logger := svc.logger.With(slog.String("op", "AddNewTransaction"), slog.String("txId", tx.ID))
	logger.Debug("going to handle new transaction")
	gateway := svc.getGateway(tx)
	if gateway == nil {
		logger.Error("there is no provider to handle the request")
		return ErrNoProviderFound
	}
}