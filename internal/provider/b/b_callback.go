package b

import (
	"context"
	"encoding/xml"
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

type BCallbackController struct {
	withdrawalService withdrawalService
	depositService depositService
}

func NewBCallbackController(
	withdrawalService withdrawalService,
	depositService depositService,
) *BCallbackController {
	return &BCallbackController{
		withdrawalService: withdrawalService,
		depositService: depositService,
	}
}

// Struct for Withdrawal Callback Data
type callbackWithdrawalData struct {
	XMLName    xml.Name `xml:"WithdrawalCallback"`
	RequestID  string   `xml:"RequestID"`
	TrackingID string   `xml:"TrackingID"`
	IsSuccess  bool     `xml:"IsSuccess"`
}

func (req *callbackWithdrawalData) toDto() dto.UpdateWithdrawalStatusDTO {
	return dto.UpdateWithdrawalStatusDTO{
		EventId:    req.RequestID,
		TrackingId: req.TrackingID,
		IsSuccess:  req.IsSuccess,
	}
}

func (ctrl BCallbackController) HandleWithdrawCallback(ctx echo.Context) error {
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

// Struct for Deposit Callback Data
type depositCallbackData struct {
	XMLName    xml.Name        `xml:"DepositCallback"`
	RequestId  string          `xml:"RequestID"`
	TrackingId string          `xml:"TrackingID"`
	IBAN       string          `xml:"IBAN"`
	Amount     decimal.Decimal `xml:"Amount"`
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

func (ctrl *BCallbackController) HandleDepositCallback(ctx echo.Context) error {
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

func (ctrl *BCallbackController) Register(gp *echo.Group) {
	group := gp.Group("/b")

	withdrawalGp := group.Group("/withdraw")
	withdrawalGp.POST("/callback", ctrl.HandleWithdrawCallback)

	depositGp := group.Group("/deposit")
	depositGp.POST("/callback", ctrl.HandleDepositCallback)
}
