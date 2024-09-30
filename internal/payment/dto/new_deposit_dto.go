package dto

import "github.com/shopspring/decimal"

type NewDepositDTO struct {
	RequestId  string
	TrackingId string
	IBAN       string
	Gateway    string
	Amount     decimal.Decimal
}
