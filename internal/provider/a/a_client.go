package a

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/alirezazeynali75/exify/internal/payment"
	"github.com/sony/gobreaker"
)

type httpClientWithCb interface {
	GetState() gobreaker.State
	Post(ctx context.Context, uri string, body string, header map[string]string) (string, error)
	Get(ctx context.Context, uri string, header map[string]string) (string, error)
}

type AProvider struct {
	logger     *slog.Logger
	httpClient httpClientWithCb
	token string
}

func NewAProvider(
	logger *slog.Logger,
	httpClient httpClientWithCb,
	token string,
) *AProvider {
	return &AProvider{
		logger: logger,
		httpClient: httpClient,
		token: token,
	}
}

func (p *AProvider) CanDo(tx payment.Withdrawal) bool {
	return p.httpClient.GetState() == gobreaker.StateOpen
}


func (p *AProvider) GetName() string {
	return "AProvider"
}

func (p *AProvider) Execute(ctx context.Context, tx payment.Withdrawal) (string, error) {
	req := aProviderCashOutRequest{
		ClientID: tx.ID,
		Destination: tx.Destination,
		Amount: tx.Amount,
	}

	reqStringify, err := json.Marshal(req)

	if err != nil {
		return "", err
	}

	resp, err := p.httpClient.Post(ctx, "/cashout", string(reqStringify), map[string]string{
		"API_TOKEN": p.token,
	})

	if err != nil {
		return "", err
	}

	var response aProviderCashOutResponse

	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return "", err
	}
	return response.TrackingId, nil
}
