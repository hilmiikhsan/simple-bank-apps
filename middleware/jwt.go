package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/simple-bank-apps/config"
	"github.com/simple-bank-apps/constants"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -source=./jwt.go -destination=./mocks/jwt/jwt.go -package=jwt_service_mocks
type JWT interface {
	GenerateJWTToken(ctx context.Context, key string, req JWTRequest) (string, error)
	ExtractJWTClaims(ctx context.Context, token string) (claims *JWTClaims, err error)
	SaveTokenToRedis(ctx context.Context, timeLimit int, token, id, authKey string) error
	ValidateTokenIssuer(claims *JWTClaims) (err error)
	GetTokenFromRedis(ctx context.Context, id, authKey string) (string, error)
	DeleteTokenFromRedis(ctx context.Context, id string, authKey string) error
}

func NewJWT(config *config.Config, redis *redis.Client) JWT {
	return &jwtObject{
		config: config,
		redis:  redis,
	}
}

type jwtObject struct {
	config *config.Config
	redis  *redis.Client
}

func (j *jwtObject) GenerateJWTToken(ctx context.Context, key string, req JWTRequest) (string, error) {
	JWTSignatureKey := []byte(j.config.JWT.Secret)
	expireTime := time.Now().Add(time.Duration(j.config.JWT.TokenLifeTimeHour) * time.Hour)

	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.config.JWT.Issuer,
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ID:       req.ID,
		Username: req.Username,
		OrigIat:  jwt.TimeFunc().Unix(),
	}

	token := jwt.NewWithClaims(
		JWTSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString(JWTSignatureKey)
	if err != nil {
		return "", err
	}

	err = j.SaveTokenToRedis(ctx, j.config.JWT.TokenLifeTimeHour, signedToken, req.ID, key)
	if err != nil {
		err = constants.ErrSaveTokenToRedis
		log.Error(err)
		return "", err
	}

	return signedToken, nil
}

func (j *jwtObject) ExtractJWTClaims(ctx context.Context, token string) (claims *JWTClaims, err error) {
	key := "auth-customer"

	splitToken := strings.Split(token, bearer)
	if len(splitToken) != 2 {
		return nil, constants.ErrTokenIsRequired
	}
	reqToken := strings.TrimSpace(splitToken[1])

	t, _ := jwt.ParseWithClaims(reqToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.config.JWT.Secret, nil
	})

	claims = t.Claims.(*JWTClaims)

	err = j.ValidateTokenIssuer(claims)
	if err != nil {
		log.WithFields(log.Fields{
			"ID":    claims.Id,
			"Error": err.Error(),
		}).Error("Someone is trying to access with different issuer token")
		return nil, err
	}

	tokenRedis, err := j.GetTokenFromRedis(ctx, claims.ID, key)
	if err != nil {
		err := constants.ErrGetTokenToRedis
		log.Print(fmt.Sprintf("%s: %s", constants.ErrGetTokenToRedis, err.Error()))
		return nil, err
	}

	if tokenRedis != reqToken {
		err := constants.ErrPleaseReLogin
		return nil, err
	}

	return claims, nil
}

func (j *jwtObject) SaveTokenToRedis(ctx context.Context, timeLimit int, token, id, authKey string) error {
	key := fmt.Sprintf("%s:%s", authKey, id)
	ttl := time.Duration(timeLimit) * time.Hour
	err := j.redis.Set(ctx, key, token, ttl).Err()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (j *jwtObject) ValidateTokenIssuer(claims *JWTClaims) (err error) {
	if claims.Issuer != j.config.JWT.Issuer {
		return err
	}

	return nil
}

func (j *jwtObject) GetTokenFromRedis(ctx context.Context, id, authKey string) (string, error) {
	key := fmt.Sprintf("%s:%s", authKey, id)
	val, err := j.redis.Get(ctx, key).Result()
	if err != nil {
		return "", nil
	}

	return val, nil
}

func (j *jwtObject) DeleteTokenFromRedis(ctx context.Context, id string, authKey string) error {
	key := fmt.Sprintf("%s:%s", authKey, id)
	_, err := j.redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
