package repo

import (
	"time"

	"github.com/alirezazeynali75/exify/internal/payment"
	"github.com/shopspring/decimal"
)

type DepositModel struct {
		ID         string          `gorm:"column:id;type:char(36);primaryKey"`     // Assuming a UUID for ID
		TrackingId string          `gorm:"column:tracking_id;type:varchar(100);not null"` // Tracking ID with varchar and not null
		IBAN       string          `gorm:"column:iban;type:varchar(34);not null"`  // IBAN with varchar length and not null constraint
		Gateway    string          `gorm:"column:gateway;type:varchar(50);not null"` // Gateway name with varchar length and not null constraint
		Amount     decimal.Decimal `gorm:"column:amount;type:decimal(20,4);not null"` // Amount with precision and not null
		CreatedAt  time.Time       `gorm:"column:created_at;type:timestamp;not null;autoCreateTime"` // Auto-set created time
		UpdatedAt  time.Time       `gorm:"column:updated_at;type:timestamp;not null;autoUpdateTime"` // Auto-set updated time
}

func (DepositModel) TableName() string {
	return "deposits"
}

func (m *DepositModel) ToDomainModel() payment.Deposit {
	return payment.Deposit{
		ID: m.ID,
		TrackingId: m.TrackingId,
		IBAN: m.IBAN,
		Gateway: m.Gateway,
		Amount: m.Amount,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FromDomainModelToDeposit(m payment.Deposit) DepositModel {
	return DepositModel{
		ID: m.ID,
		TrackingId: m.TrackingId,
		IBAN: m.IBAN,
		Gateway: m.Gateway,
		Amount: m.Amount,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
} 