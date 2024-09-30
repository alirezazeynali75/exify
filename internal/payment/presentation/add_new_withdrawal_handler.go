package presentation

import (
	"net/http"

	"github.com/alirezazeynali75/exify/internal/payment/dto"
	"github.com/alirezazeynali75/exify/pkg/responses"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type addNewWithdrawalRequest struct {
	RequestID   string          `json:"request_id"`
	ID          string          `json:"id"`
	Amount      decimal.Decimal `json:"amount"`
	Destination string          `json:"destination"`
}

func (req *addNewWithdrawalRequest) ToDto() dto.NewWithdrawalDto {
	return dto.NewWithdrawalDto{
		EventId:     req.RequestID,
		ID:          req.ID,
		Amount:      req.Amount,
		Destination: req.Destination,
	}
}

func (ctrl *PaymentController) AddNewWithdrawalRequestHandler(ctx echo.Context) error {
	var req addNewWithdrawalRequest

	err := ctx.Bind(&req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewErrorResponse("validation failed"))
	}

	err = ctrl.withdrawalSvc.AddNewWithdrawTransaction(ctx.Request().Context(), req.ToDto())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewErrorResponse("internal error occurred"))
	}
	return ctx.JSON(http.StatusOK, responses.NewResponse("transaction created", map[string]string{}))
}
