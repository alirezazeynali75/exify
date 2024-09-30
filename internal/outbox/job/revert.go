package job

import (
	"context"
	"log/slog"

	"github.com/robfig/cron"
)

type RevertStalledMessageJob struct {
	logger *slog.Logger
	svc    outboxService
}

func NewRevertStalledMessageJob(
	logger *slog.Logger,
	svc outboxService,
) *RevertStalledMessageJob {
	return &RevertStalledMessageJob{
		logger: logger,
		svc:    svc,
	}
}

func (job *RevertStalledMessageJob) Run() {
	err := job.svc.RevertPending(context.Background())
	if err != nil {
		job.logger.Error("there is an error")
		return
	}
	job.logger.Info("the message reverted")
}

func (job *RevertStalledMessageJob) Register(c *cron.Cron) {
	c.AddJob("@every 10m", job)
}
