package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/simple-bank-apps/constants"
	"github.com/simple-bank-apps/utils"
)

var tokenCtxKey = &contextKey{"token"}

var TokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}

type Token struct {
	Value string
}

type TokenMiddleware interface {
	TokenAuthorize() gin.HandlerFunc
	TokenMiddlewareAuthorize() gin.HandlerFunc
}

func NewTokenMiddleware(jwt JWT) TokenMiddleware {
	return &tokenMiddleware{
		jwt,
	}
}

type tokenMiddleware struct {
	JWTService JWT
}

func (t *tokenMiddleware) TokenAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		err := constants.ErrTokenIsRequired
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.NewUnauthorized(err.Error()))
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), tokenCtxKey, &Token{Value: tokenHeader})
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (t *tokenMiddleware) TokenMiddlewareAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		err := constants.ErrTokenIsRequired
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.NewUnauthorized(err.Error()))
			c.Abort()
			return
		}

		claims, err := t.JWTService.ExtractJWTClaims(c.Request.Context(), tokenHeader)
		if err != nil {
			if strings.Contains(err.Error(), constants.ErrKeyIsNotInvalidType.Error()) || strings.Contains(err.Error(), constants.ErrGetTokenToRedis.Error()) || strings.Contains(err.Error(), constants.ErrPleaseReLogin.Error()) || strings.Contains(err.Error(), constants.ErrSaveTokenToRedis.Error()) {
				c.JSON(http.StatusUnauthorized, utils.NewUnauthorized(err.Error()))
				c.Abort()
				return
			}

			c.JSON(http.StatusInternalServerError, utils.NewInternalServerError(err.Error()))
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), tokenCtxKey, claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetTokenFromContext(ctx context.Context) *Token {
	raw, _ := ctx.Value(tokenCtxKey).(*Token)
	return raw
}

func GetTokenMiddlewareFromContext(ctx context.Context) *JWTClaims {
	raw, _ := ctx.Value(tokenCtxKey).(*JWTClaims)
	return raw
}
