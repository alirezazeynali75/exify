package presentation

import (
	"context"

	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/labstack/echo/v4"
)

type withdrawalService interface {
	AddNewWithdrawTransaction(ctx context.Context, d dto.NewWithdrawalDto) error
}

type PaymentController struct {
	withdrawalSvc withdrawalService
}

func NewPaymentController(
	withdrawalService withdrawalService,
) *PaymentController {
	return &PaymentController{
		withdrawalSvc: withdrawalService,
	}
}

func (ctrl *PaymentController) Register(gp *echo.Group) {
	group := gp.Group("/payment")

	group.POST("/cashout", ctrl.AddNewWithdrawalRequestHandler)
}