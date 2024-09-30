package dto

type NewDepositDTO struct {
	RequestId  string
	TrackingId string
	IBAN       string
	Gateway    string
	Amount     string
}
