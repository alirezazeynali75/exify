package config

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/sony/gobreaker"
)

type AProvider struct {
	ClientURL string `env:"URL" envDefault:"http://0.0.0.0:8080"`
	Token     string `env:"Token" envDefault:""`
}


func (p *AProvider) GetCircuitBreakerConfig() *gobreaker.CircuitBreaker {
	onStateChangeCb := func(name string, from, to gobreaker.State) {
		slog.Info(fmt.Sprintf("CircuitBreaker '%s' changed from '%s' to '%s'", name, from, to))
	}
	readyToTripCb := func(counts gobreaker.Counts) bool {
		return counts.ConsecutiveFailures > uint32(5)
	}
	return gobreaker.NewCircuitBreaker(
		gobreaker.Settings{
			Name:         	"AProvider",
			MaxRequests:   uint32(5),
			Timeout:       time.Second * 10,
			Interval:      time.Second * 2,
			ReadyToTrip:   readyToTripCb,
			OnStateChange: onStateChangeCb,
		},
	)
}