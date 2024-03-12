package middleware

import (
	"crud/internal/core/model"
	"crud/internal/core/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {

	claims := validateAuth(c)

	c.Set("user", claims.Login)

	c.Next()
}

func parseToken(token string) (service.TokenClaims, error) {
	// at - access token
	at, err := jwt.ParseWithClaims(token, &service.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid method")
			}

			return []byte(model.SignInKey), nil
		})

	if err != nil {
		return service.TokenClaims{}, err
	}

	claims, ok := at.Claims.(*service.TokenClaims)

	if !ok {
		return service.TokenClaims{}, err
	}

	return *claims, nil
}

func validateAuth(c *gin.Context) service.TokenClaims {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid")
		return service.TokenClaims{}
	}

	splitted := strings.Split(auth, " ")

	claims, err := parseToken(splitted[1])

	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
	}

	return claims
}

func AdminAuthMiddleware(c *gin.Context) {
	claims := validateAuth(c)
	if !claims.IsAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid")
		return
	}
	c.Set("user", claims.Login)

	c.Next()
}
