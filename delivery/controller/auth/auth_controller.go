package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/simple-bank-apps/constants"
	"github.com/simple-bank-apps/dto"
	"github.com/simple-bank-apps/middleware"
	"github.com/simple-bank-apps/usecase/auth"
	"github.com/simple-bank-apps/utils"
	"github.com/simple-bank-apps/validation"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	router      *gin.Engine
	authUsecase auth.AuthUsecase
}

func (a *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest(validation.FieldErrors(err)))
		return
	}

	err := a.authUsecase.Register(c.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUsernameAlreadyExist.Error()) || strings.Contains(err.Error(), constants.ErrAccountNumberAlreadyExist.Error()) {
			c.JSON(http.StatusBadRequest, utils.NewBadRequest(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.NewSuccessResponse(http.StatusCreated, "Success Register Customer"))
}

func (a *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewBadRequest(validation.FieldErrors(err)))
		return
	}

	response, err := a.authUsecase.Login(c.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUsernameOrPasswordNotMatch.Error()) {
			c.JSON(http.StatusBadRequest, utils.NewUnauthorized(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponseWithData(http.StatusOK, response, "Success Login Customer"))
}

func (a *AuthController) Logout(c *gin.Context) {
	token := middleware.GetTokenFromContext(c.Request.Context())

	err := a.authUsecase.Logout(c, token)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrInvalidToken.Error()) {
			c.JSON(http.StatusBadRequest, utils.NewBadRequest(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse(http.StatusOK, "Success Logout Customer"))
}

func NewAuthController(
	router *gin.Engine,
	authUsecase auth.AuthUsecase,
	tokenMiddleware middleware.TokenMiddleware,
	logger middleware.LogMiddleware,
) {
	controller := &AuthController{
		router:      router,
		authUsecase: authUsecase,
	}

	routerGroup := router.Group("/api/v1")
	routerGroup.Use(logger.LogRequestAuth(logrus.New()))

	routerGroup.POST("/register", controller.Register)
	routerGroup.POST("/login", controller.Login)
	routerGroup.POST("/logout", tokenMiddleware.TokenAuthorize(), controller.Logout)
}
