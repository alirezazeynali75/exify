package repo

import (
	"database/sql"
	"time"

	sqlutil "github.com/alirezazeynali75/exify/pkg/sql"

	"github.com/alirezazeynali75/exify/internal/payment"
	"github.com/shopspring/decimal"
)

type WithdrawalModel struct {
	ID          string          `gorm:"primaryKey;column:id"`
	TrackingId  sql.NullString  `gorm:"column:tracking_id"`
	Destination string          `gorm:"column:destination"`
	Gateway     string          `gorm:"column:gateway"`
	Amount      decimal.Decimal `gorm:"column:amount"`
	Status      string          `gorm:"column:status"`
	CreatedAt   time.Time       `gorm:"column:created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at"`
}

func (WithdrawalModel) TableName() string {
	return "withdrawals"
}

func (m WithdrawalModel) ToDomainModel() payment.Withdrawal {
	return payment.Withdrawal{
		ID:          m.ID,
		TrackingId:  *sqlutil.ToPtrString(m.TrackingId),
		Destination: m.Destination,
		Gateway:     m.Gateway,
		Amount:      m.Amount,
		Status:      payment.PaymentStatus(m.Status),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func FromDomainModel(d payment.Withdrawal) WithdrawalModel {
	return WithdrawalModel{
		ID:          d.ID,
		TrackingId:  sqlutil.ToNullableString(&d.TrackingId),
		Destination: d.Destination,
		Gateway:     d.Gateway,
		Amount:      d.Amount,
		Status:      d.Status.String(),
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
