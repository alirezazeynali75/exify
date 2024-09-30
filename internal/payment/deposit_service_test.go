package payment

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock definitions

type mockInboxRepo struct {
	mock.Mock
}

func (m *mockInboxRepo) InsertEvent(ctx context.Context, requestId string) error {
	args := m.Called(ctx, requestId)
	return args.Error(0)
}

type mockDepositRepo struct {
	mock.Mock
}

func (m *mockDepositRepo) CreateNewTransaction(ctx context.Context, deposit Deposit) error {
	args := m.Called(ctx, deposit)
	return args.Error(0)
}

type mockOutboxRepo struct {
	mock.Mock
}

func (m *mockOutboxRepo) InsertNewEvent(ctx context.Context, event string, topic string) error {
	args := m.Called(ctx, event, topic)
	return args.Error(0)
}

type mockSession struct {
	mock.Mock
}

// Begin implements db.Session.
func (m *mockSession) Begin(ctx context.Context) (db.Session, error) {
	panic("unimplemented")
}

// Commit implements db.Session.
func (m *mockSession) Commit() error {
	panic("unimplemented")
}

// Context implements db.Session.
func (m *mockSession) Context() context.Context {
	panic("unimplemented")
}

// Rollback implements db.Session.
func (m *mockSession) Rollback() error {
	panic("unimplemented")
}

func (m *mockSession) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	_ = m.Called(ctx, fn)
	return fn(ctx)
}

// Test case

func TestDepositService_AddDeposit(t *testing.T) {
	tests := []struct {
		name             string
		input            dto.NewDepositDTO
		mockInboxError   error
		mockDepositError error
		mockOutboxError  error
		expectedError    error
		mockTxError      error
	}{
		{
			name:             "successful deposit",
			input:            dto.NewDepositDTO{RequestId: "req1"},
			mockInboxError:   nil,
			mockDepositError: nil,
			mockOutboxError:  nil,
			expectedError:    nil,
		},
		{
			name:           "inbox insertion error",
			input:          dto.NewDepositDTO{RequestId: "req2"},
			mockInboxError: errors.New("inbox error"),
			expectedError:  errors.New("inbox error"),
		},
		{
			name:             "deposit creation error",
			input:            dto.NewDepositDTO{RequestId: "req3"},
			mockInboxError:   nil,
			mockDepositError: errors.New("deposit error"),
			expectedError:    errors.New("deposit error"),
		},
		{
			name:             "outbox insertion error",
			input:            dto.NewDepositDTO{RequestId: "req4"},
			mockInboxError:   nil,
			mockDepositError: nil,
			mockOutboxError:  errors.New("outbox error"),
			expectedError:    errors.New("outbox error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock objects
			mockInbox := new(mockInboxRepo)
			mockDeposit := new(mockDepositRepo)
			mockOutbox := new(mockOutboxRepo)
			mockSession := new(mockSession)

			// Initialize logger (could be mocked or set to nil)
			logger := slog.New(slog.NewTextHandler(io.Discard, nil))

			// Create the service
			service := NewDepositService(logger, mockSession, mockInbox, mockDeposit, mockOutbox)

			// Set up expectations
			mockInbox.On("InsertEvent", mock.Anything, tt.input.RequestId).Return(tt.mockInboxError)
			mockDeposit.On("CreateNewTransaction", mock.Anything, mock.Anything).Return(tt.mockDepositError)
			mockOutbox.On("InsertNewEvent", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockOutboxError)
			mockSession.On("Transaction", mock.Anything, mock.Anything).Return(tt.mockTxError)

			// Execute the function being tested
			err := service.AddDeposit(context.Background(), tt.input)

			// Check the results
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
