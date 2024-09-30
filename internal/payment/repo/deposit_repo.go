package repo

import (
	"context"

	"github.com/alirezazeynali75/exify/internal/db"
	"github.com/alirezazeynali75/exify/internal/payment"
	"gorm.io/gorm"
)


type DepositRepo struct {
	db *gorm.DB
}

func NewDepositRepo(
	db *gorm.DB,
) *DepositRepo {
	return &DepositRepo{
		db: db,
	}
}

func (repo *DepositRepo) CreateNewTransaction(ctx context.Context, tx payment.Deposit) error {
	model := FromDomainModelToDeposit(tx)


	return db.DB(ctx, repo.db).Model(&DepositModel{}).WithContext(ctx).Create(&model).Error
}