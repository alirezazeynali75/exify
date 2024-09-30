package payment

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID        string
	Type      PaymentType
	Gateway   string
	Amount    decimal.Decimal
	Status    PaymentStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
