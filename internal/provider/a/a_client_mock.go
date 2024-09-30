package a

import (
	"context"

	"github.com/alirezazeynali75/exify/internal/payment"
	"github.com/google/uuid"
)

type AProviderMock struct {
}

func NewAProviderMock() *AProviderMock {
	return &AProviderMock{}
}

func (p *AProviderMock) CanDo(tx payment.Withdrawal) bool {
	return true
}

func (p *AProviderMock) GetName() string {
	return "AProvider"
}

func (p *AProviderMock) Execute(ctx context.Context, tx payment.Withdrawal) (string, error) {
	return uuid.NewString(), nil
}