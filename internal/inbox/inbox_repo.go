package inbox

import (
	"context"

	"github.com/alirezazeynali75/exify/internal/db"
	"gorm.io/gorm"
)

type InboxRepo struct {
	db *gorm.DB
}

func NewInboxRepo(
	db *gorm.DB,
) *InboxRepo {
	return &InboxRepo{
		db: db,
	}
}

func (repo *InboxRepo) InsertEvent(ctx context.Context, eventId string) error {
	return db.DB(ctx, repo.db).Model(&InboxModel{}).WithContext(ctx).Create(&InboxModel{
		EventID: eventId,
	}).Error
}