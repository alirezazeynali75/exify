package a

import (
	"context"
	"net/http"

	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/alirezazeynali75/exify/pkg/responses"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type withdrawalService interface {
	UpdateWithdrawalStatus(ctx context.Context, d dto.UpdateWithdrawalStatusDTO) error
}

type depositService interface {
	AddDeposit(ctx context.Context, d dto.NewDepositDTO) error
}

type ACallbackController struct {
	withdrawalService withdrawalService
	depositService depositService
}

func NewACallbackController(
	withdrawalService withdrawalService,
	depositService depositService,
) *ACallbackController {
	return &ACallbackController{
		withdrawalService: withdrawalService,
		depositService: depositService,
	}
}

type callbackWithdrawalData struct {
	RequestID  string `json:"request_id"`
	TrackingID string `json:"tracking_id"`
	IsSuccess  bool   `json:"is_success"`
}

func (req *callbackWithdrawalData) toDto() dto.UpdateWithdrawalStatusDTO {
	return dto.UpdateWithdrawalStatusDTO{
		EventId:    req.RequestID,
		TrackingId: req.TrackingID,
		IsSuccess:  req.IsSuccess,
	}
}

func (ctrl ACallbackController) HandleWithdrawCallback(ctx echo.Context) error {
	var req callbackWithdrawalData

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewErrorResponse("validation failed"))
	}

	err = ctrl.withdrawalService.UpdateWithdrawalStatus(ctx.Request().Context(), req.toDto())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewErrorResponse("internal error occurred"))
	}
	return ctx.JSON(http.StatusOK, responses.NewResponse("transaction updated", map[string]string{}))
}

type depositCallbackData struct {
	RequestId  string          `json:"request_id"`
	TrackingId string          `json:"tracking_id"`
	IBAN       string          `json:"IBAN"`
	Amount     decimal.Decimal `json:"amount"`
}

func (req *depositCallbackData) toDto() dto.NewDepositDTO {
	return dto.NewDepositDTO{
		RequestId: req.RequestId,
		TrackingId: req.TrackingId,
		IBAN: req.IBAN,
		Amount: req.Amount,
		Gateway: "A",
	}
}

func (ctrl *ACallbackController) HandleDepositCallback(ctx echo.Context) error {
	var req depositCallbackData

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewErrorResponse("validation failed"))
	}

	err = ctrl.depositService.AddDeposit(ctx.Request().Context(), req.toDto())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewErrorResponse("internal error occurred"))
	}
	return ctx.JSON(http.StatusOK, responses.NewResponse("deposit updated", map[string]string{}))
}

func (ctrl *ACallbackController) Register(gp *echo.Group) {
	group := gp.Group("/a")

	withdrawalGp := group.Group("/withdraw")
	withdrawalGp.POST("/callback", ctrl.HandleWithdrawCallback)

	depositGp := group.Group("/deposit")
	depositGp.POST("/callback", ctrl.HandleDepositCallback)
}
