package b

import (
	"context"

	"github.com/alirezazeynali75/exify/internal/payment"
	"github.com/google/uuid"
)

type BProviderMock struct {
}

func NewBProviderMock() *BProviderMock {
	return &BProviderMock{}
}

func (p *BProviderMock) CanDo(tx payment.Withdrawal) bool {
	return true
}

func (p *BProviderMock) GetName() string {
	return "BProvider"
}

func (p *BProviderMock) Execute(ctx context.Context, tx payment.Withdrawal) (string, error) {
	return uuid.NewString(), nil
}