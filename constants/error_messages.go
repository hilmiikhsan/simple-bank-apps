package constants

import "errors"

var (
	ErrUsernameAlreadyExist       = errors.New("username already exist")
	ErrSaveTokenToRedis           = errors.New("cannot save token to redis")
	ErrTokenIsRequired            = errors.New("token is required")
	ErrKeyIsNotInvalidType        = errors.New("key is not invalid type")
	ErrGetTokenToRedis            = errors.New("cannot get token from redis")
	ErrPleaseReLogin              = errors.New("please re-login")
	ErrUsernameOrPasswordNotMatch = errors.New("username or password not match")
	ErrInvalidToken               = errors.New("invalid token")
)
