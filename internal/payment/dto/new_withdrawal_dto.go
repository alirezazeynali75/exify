package dto

import (
	"github.com/shopspring/decimal"
)

type NewWithdrawalDto struct {
	EventId string
	ID      string
	Amount  decimal.Decimal
	Destination string
}
