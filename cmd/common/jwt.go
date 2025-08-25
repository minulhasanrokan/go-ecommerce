package common

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
)

type CustomJWTClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.UserModel) (*string, *string, error) {

	userClaims := CustomJWTClaims{
		ID: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour + 100)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signedAccessToken, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {

		return nil, nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomJWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour + 100)),
		},
	})

	signedRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {

		return nil, nil, err
	}

	return &signedAccessToken, &signedRefreshToken, nil
}

func ParseJWTSignedAccessToken(token string) (*CustomJWTClaims, error) {

	parsedJWTAccessToken, err := jwt.ParseWithClaims(token, &CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := parsedJWTAccessToken.Claims.(*CustomJWTClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("unknown claims type, cannot parse jwt token")
	}
}

func IsClaimExpire(claims *CustomJWTClaims) bool {

	currentTime := jwt.NewNumericDate(time.Now())

	return claims.ExpiresAt.Time.Before(currentTime.Time)
}
