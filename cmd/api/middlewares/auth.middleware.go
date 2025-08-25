package middlewares

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/common"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
	"gorm.io/gorm"
)

func (appMiddleware *AppMiddleware) AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		c.Response().Header().Set("Vary", "Authorization")
		authHeader := c.Request().Header.Get("Authorization")

		if strings.HasPrefix(authHeader, "Bearer ") == false {

			return common.ApiSendUnauthorizedResponse(c, "Please provide a Bearer token")
		}

		authHeaderSplit := strings.Split(authHeader, " ")

		accessToken := authHeaderSplit[1]

		claims, err := common.ParseJWTSignedAccessToken(accessToken)

		if err != nil {
			return common.ApiSendUnauthorizedResponse(c, err.Error())
		}

		if common.IsClaimExpire(claims) == true {

			return common.ApiSendUnauthorizedResponse(c, "Token is expired")
		}

		var user models.UserModel

		result := appMiddleware.Db.First(&user, claims.ID)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return common.ApiSendUnauthorizedResponse(c, "Invalid access token")
		}

		if result.Error != nil {
			return common.ApiSendUnauthorizedResponse(c, "Invalid access token")
		}

		c.Set("user", user)

		_, ok := c.Get("user").(models.UserModel)

		if !ok {
			return common.ApiSendUnauthorizedResponse(c, "User authentication failed")
		}

		return next(c)
	}
}
