package middlewares

import (
	"api-alta-dashboard/config"
	"api-alta-dashboard/utils/helper"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key string

func InitJWT(c *config.AppConfig) {
	key = c.JWT_SECRET
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(key),
	})
}

func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))

}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}

func ExtractTokenUserRole(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		return role
	}
	return ""
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		user := e.Get("user").(*jwt.Token)
		if user.Valid {
			claims := user.Claims.(jwt.MapClaims)
			role := claims["role"].(string)

			if role != "Super Admin" {
				return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Error. User role not authorized to access."))
			}
		}
		return next(e)

	}
}
