package b

import (
	"context"
	"encoding/xml"
	"log/slog"

	"github.com/alirezazeynali75/exify/internal/payment"
	"github.com/sony/gobreaker"
)

type httpClientWithCb interface {
	GetState() gobreaker.State
	Post(ctx context.Context, uri string, body string, header map[string]string) (string, error)
	Get(ctx context.Context, uri string, header map[string]string) (string, error)
}

type BProvider struct {
	logger     *slog.Logger
	httpClient httpClientWithCb
	token string
}

func NewBProvider(
	logger *slog.Logger,
	httpClient httpClientWithCb,
	token string,
) *BProvider {
	return &BProvider{
		logger: logger,
		httpClient: httpClient,
		token: token,
	}
}

func (p *BProvider) CanDo(tx payment.Withdrawal) bool {
	return p.httpClient.GetState() == gobreaker.StateOpen
}


func (p *BProvider) GetName() string {
	return "BProvider"
}

func (p *BProvider) Execute(ctx context.Context, tx payment.Withdrawal) (string, error) {
	req := bProviderCashOutRequest{
		ClientID: tx.ID,
		Destination: tx.Destination,
		Amount: tx.Amount,
	}

	reqStringify, err := xml.Marshal(req)

	if err != nil {
		return "", err
	}

	resp, err := p.httpClient.Post(ctx, "/cashout", string(reqStringify), map[string]string{
		"API_TOKEN": p.token,
	})

	if err != nil {
		return "", err
	}

	var response bProviderCashOutResponse

	err = xml.Unmarshal([]byte(resp), &response)
	if err != nil {
		return "", err
	}
	return response.TrackingID, nil
}
