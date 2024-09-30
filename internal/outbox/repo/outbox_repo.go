package outbox

import (
	"context"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/outbox"
	"gorm.io/gorm"
)

type OutboxRepo struct {
	db *gorm.DB
}


func (repo *OutboxRepo) InsertNewEvent(ctx context.Context, event string, topic string) error {
	model := OutboxModel{
		Payload: event,
		Topic: topic,
		Status: outbox.READY.String(),
	}
	return db.DB(ctx, repo.db).Model(&OutboxModel{}).WithContext(ctx).Create(&model).Error
}