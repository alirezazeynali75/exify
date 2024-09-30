package a

import "github.com/shopspring/decimal"

type aProviderCashOutRequest struct {
	ClientID    string          `json:"client_id"`
	Destination string          `json:"destination"`
	Amount      decimal.Decimal `json:"amount"`
}

type aProviderCashOutResponse struct {
	TrackingId string `json:"tracking_id"`
}
