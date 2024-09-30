package repo

import (
	"context"
	"time"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/outbox"
	"gorm.io/gorm"
)

type OutboxRepo struct {
	db *gorm.DB
}

func NewOutboxRepo(
	db *gorm.DB,
) *OutboxRepo {
	return &OutboxRepo{
		db: db,
	}
}

func (repo *OutboxRepo) InsertNewEvent(ctx context.Context, event string, topic string) error {
	model := OutboxModel{
		Payload: event,
		Topic: topic,
		Status: outbox.READY.String(),
	}
	return db.DB(ctx, repo.db).Model(&OutboxModel{}).WithContext(ctx).Create(&model).Error
}


func (repo *OutboxRepo) GetPendingEventsAndUpdateStatus(ctx context.Context) ([]outbox.Outbox, error) {
	var records []OutboxModel
	err := db.DB(ctx, repo.db).Model(&OutboxModel{}).WithContext(ctx).Where("status = ?", outbox.READY).Find(&records).Error
	if err != nil {
		return []outbox.Outbox{}, err
	}

	domainRecords := make([]outbox.Outbox, len(records))

	for i, record := range records {
		domainRecords[i] = record.ToDomainModel()
	}

	return domainRecords, nil
}

func (repo *OutboxRepo) UpdateByID(ctx context.Context, id uint64, status outbox.OutboxStatus) error {
	return db.DB(ctx, repo.db).Model(&OutboxModel{}).WithContext(ctx).Where("id = ?", id).Updates(map[string]interface{}{
		"status": status,
	}).Error
}

func (repo *OutboxRepo) BatchUpdatePendingBasedOnTime(ctx context.Context, status outbox.OutboxStatus, maxAge time.Duration) error {
	now := time.Now()

	passedTime := now.Add(-maxAge)

	err := db.DB(ctx, repo.db).Model(&OutboxModel{}).WithContext(ctx).Where("updated_at <= ?", passedTime).Updates(map[string]interface{}{
		"status": status,
	}).Error

	return err
}
