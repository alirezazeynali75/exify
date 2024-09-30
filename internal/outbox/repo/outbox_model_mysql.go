package repo

import (
	"time"

	"github.com/alirezazeynali75/exify/internal/outbox"
)

type OutboxModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Payload   string    `gorm:"type:text;not null"`
	Topic     string    `gorm:"type:varchar(255);not null"`
	Status    string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (m OutboxModel) ToDomainModel() outbox.Outbox {
	return outbox.Outbox{
		ID: m.ID,
		Payload: m.Payload,
		Topic: m.Topic,
		Status: m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
