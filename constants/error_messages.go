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
	ErrAccountNumberAlreadyExist  = errors.New("account number already exist")
	ErrAccountNumberNotFound      = errors.New("account number not found")
	ErrCustomerNotFound           = errors.New("customer not found")
	ErrAmountNotEnough            = errors.New("amount not enough")
)
