package payment

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/simple-bank-apps/constants"
	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/middleware"
	"github.com/simple-bank-apps/usecase/payment"
	"github.com/simple-bank-apps/utils"
	"github.com/simple-bank-apps/validation"
	"github.com/sirupsen/logrus"
)

type PaymentController struct {
	router         *gin.Engine
	paymentUsecase payment.PaymentUsecase
}

func (p *PaymentController) Payment(c *gin.Context) {
	var req dto.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest(validation.FieldErrors(err)))
		return
	}

	err := p.paymentUsecase.Payment(c.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrCustomerNotFound.Error()) || strings.Contains(err.Error(), constants.ErrAccountNumberNotFound.Error()) {
			c.JSON(http.StatusNotFound, utils.NewNotFound(err.Error()))
			return
		}

		if strings.Contains(err.Error(), constants.ErrAmountNotEnough.Error()) {
			c.JSON(http.StatusBadRequest, utils.NewBadRequest(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(http.StatusOK, "Success Payment"))
}

func NewPaymentController(
	router *gin.Engine,
	paymentUsecase payment.PaymentUsecase,
	tokenMiddleware middleware.TokenMiddleware,
	logger middleware.LogMiddleware,
) {
	controller := &PaymentController{
		router:         router,
		paymentUsecase: paymentUsecase,
	}

	routerGroup := router.Group("/api/v1/payment")
	routerGroup.Use(tokenMiddleware.TokenMiddlewareAuthorize())
	routerGroup.Use(logger.LogRequestMiddleware(logrus.New()))

	routerGroup.POST("/", controller.Payment)
}
