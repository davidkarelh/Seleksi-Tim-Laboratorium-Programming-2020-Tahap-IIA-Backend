package middlewares

import (
	"log"
	"net/http"

	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/helper"
	"github.com/davidkarelh/Seleksi-Tim-Laboratorium-Programming-2020-Tahap-IIA-Backend/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtS services.IJWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if len(authHeader) <= len(BEARER_SCHEMA) {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := services.NewJWTService().ValidateToken(tokenString)

		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println(claims)
			log.Println("Claims[username]", claims["username"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", "Invalid token or signing method", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
