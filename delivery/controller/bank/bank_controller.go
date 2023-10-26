package bank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simple-bank-apps/middleware"
	"github.com/simple-bank-apps/usecase/bank"
	"github.com/simple-bank-apps/utils"
)

type BankController struct {
	router      *gin.Engine
	bankUsecase bank.BankUsecase
}

func (b *BankController) GetListBank(c *gin.Context) {
	data, err := b.bankUsecase.GetListBank(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponseWithData(http.StatusOK, data, "Success Get List Bank"))
}

func NewBankController(
	router *gin.Engine,
	bankUsecase bank.BankUsecase,
	tokenMiddleware middleware.TokenMiddleware,
	logger middleware.LogMiddleware,
) {
	controller := &BankController{
		router:      router,
		bankUsecase: bankUsecase,
	}

	routerGroup := router.Group("/api/v1/bank")
	routerGroup.Use(tokenMiddleware.TokenMiddlewareAuthorize())

	routerGroup.GET("/", controller.GetListBank)
}
