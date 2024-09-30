package b

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
)

type bProviderCashOutRequest struct {
	XMLName     xml.Name        `xml:"CashOutRequest"`
	ClientID    string          `xml:"ClientID"`
	Destination string          `xml:"Destination"`
	Amount      decimal.Decimal `xml:"Amount"`
}

type bProviderCashOutResponse struct {
	XMLName    xml.Name `xml:"CashOutResponse"` // Root element name
	TrackingID string   `xml:"TrackingID"`      // Element name in XML
}