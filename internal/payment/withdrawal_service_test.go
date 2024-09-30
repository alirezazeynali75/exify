package payment

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type mockWithdrawalRepo struct {
	mock.Mock
}

func (m *mockWithdrawalRepo) CreateNewTransaction(ctx context.Context, tx Withdrawal) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func (m *mockWithdrawalRepo) UpdateTrackingId(ctx context.Context, txID, trackingID string) error {
	args := m.Called(ctx, txID, trackingID)
	return args.Error(0)
}

func (m *mockWithdrawalRepo) UpdateStatusByTrackingId(ctx context.Context, trackingID string, status PaymentStatus) error {
	args := m.Called(ctx, trackingID, status)
	return args.Error(0)
}

type mockPaymentGateway struct {
	mock.Mock
}

func (m *mockPaymentGateway) CanDo(tx Withdrawal) bool {
	args := m.Called(tx)
	return args.Bool(0)
}

func (m *mockPaymentGateway) GetName() string {
	args := m.Called()
	return args.String(0)
}

func (m *mockPaymentGateway) Execute(ctx context.Context, tx Withdrawal) (string, error) {
	args := m.Called(ctx, tx)
	return args.String(0), args.Error(1)
}

// Test for AddNewWithdrawTransaction

func TestWithdrawalService_AddNewWithdrawTransaction(t *testing.T) {
	tests := []struct {
		name                string
		input               dto.NewWithdrawalDto
		gatewayCanDo        bool
		gatewayName         string
		gatewayExecuteID    string
		gatewayExecuteError error
		mockInboxError      error
		mockTxCreateError   error
		mockTxUpdateError   error
		mockOutboxError     error
		expectedError       error
	}{
		{
			name:                "successful withdrawal transaction",
			input:               dto.NewWithdrawalDto{ID: "tx1", EventId: "event1"},
			gatewayCanDo:        true,
			gatewayName:         "MockGateway",
			gatewayExecuteID:    "tracking-123",
			expectedError:       nil,
		},
		{
			name:                "no payment gateway available",
			input:               dto.NewWithdrawalDto{ID: "tx2", EventId: "event2"},
			gatewayCanDo:        false,
			expectedError:       ErrNoProviderFound,
		},
		{
			name:                "gateway execution error",
			input:               dto.NewWithdrawalDto{ID: "tx4", EventId: "event4"},
			gatewayCanDo:        true,
			gatewayName:         "MockGateway",
			gatewayExecuteID:    "tracking-123",
			gatewayExecuteError: errors.New("gateway execution error"),
			expectedError:       errors.New("gateway execution error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockInbox := new(mockInboxRepo)
			mockTxRepo := new(mockWithdrawalRepo)
			mockOutbox := new(mockOutboxRepo)
			mockSession := new(mockSession)
			mockGateway := new(mockPaymentGateway)

			logger := slog.New(slog.NewJSONHandler(io.Discard, nil))

			// Create service
			service := NewWithdrawalService(logger, mockSession, mockOutbox, mockInbox, mockTxRepo, mockGateway)

			// Set expectations
			mockGateway.On("CanDo", mock.Anything).Return(tt.gatewayCanDo)
			mockGateway.On("GetName").Return(tt.gatewayName)
			mockGateway.On("Execute", mock.Anything, mock.Anything).Return(tt.gatewayExecuteID, tt.gatewayExecuteError)
			mockInbox.On("InsertEvent", mock.Anything, mock.Anything).Return(tt.mockInboxError)
			mockTxRepo.On("CreateNewTransaction", mock.Anything, mock.Anything).Return(tt.mockTxCreateError)
			mockTxRepo.On("UpdateTrackingId", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockTxUpdateError)
			mockOutbox.On("InsertNewEvent", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockOutboxError)
			mockSession.On("Transaction", mock.Anything, mock.Anything).Return(nil)

			// Execute function under test
			err := service.AddNewWithdrawTransaction(context.Background(), tt.input)

			// Assert expectations
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
