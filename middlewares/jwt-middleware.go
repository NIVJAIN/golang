// how to use as forgetting most of the time
// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"log"
	"net/http"

	"github.com/nivjain/7-ginInterfaceMongoDBRabbitMQ-JWT-RTL/services"

	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerschema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(bearerschema):]

		// token, err := service.NewJWTService().ValidateToken(tokenString)
		token, err := services.DecodeToken(tokenString)
		// if token.Valid {
		// claims := token.Claims.(jwt.MapClaims)
		// log.Println("Claims[Name]: ", claims["name"])
		// log.Println("Claims[Admin]: ", claims["admin"])
		// log.Println("Claims[Issuer]: ", claims["iss"])
		// log.Println("Claims[IssuedAt]: ", claims["iat"])
		// log.Println("Claims[ExpiresAt]: ", claims["exp"])
		// }
		log.Println("token vis valid", token)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
