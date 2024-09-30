package repo

import (
	"context"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/payment"
	"gorm.io/gorm"
)

type WithdrawalRepo struct {
	db *gorm.DB
}

func NewWithdrawalRepo(
	db *gorm.DB,
) *WithdrawalRepo {
	return &WithdrawalRepo{
		db: db,
	}
}

func (repo *WithdrawalRepo) CreateNewTransaction(ctx context.Context, tx payment.Withdrawal) error {
	model := FromDomainModel(tx)
	err := db.DB(ctx, repo.db).Model(&WithdrawalModel{}).WithContext(ctx).Create(&model).Error
	return err
}

func (repo *WithdrawalRepo) UpdateTrackingId(ctx context.Context, txId string, trackingId string) error {
	return db.DB(ctx, repo.db).Model(&WithdrawalModel{}).WithContext(ctx).Where("id = ?", txId).Updates(map[string]interface{}{
		"tracking_id": trackingId,
	}).Error
}

func (repo *WithdrawalRepo) UpdateStatusByTrackingId(ctx context.Context, trackingId string, status payment.PaymentStatus) error {
	return db.DB(ctx, repo.db).Model(&WithdrawalModel{}).WithContext(ctx).Where("tracking_id = ?", trackingId).Updates(map[string]interface{}{
		"status": status.String(),
	}).Error
}
