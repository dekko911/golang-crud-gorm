package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/what-crud/initializers"
	"github.com/what-crud/utils"
)

func AuthMiddlewareJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("Authorization")

		// tokenString := strings.Split(tokenHeader, "Bearer ")[1]
		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		if tokenString == "" {
			ctx.JSON(utils.UNAUTH, gin.H{
				"code":  utils.UNAUTH,
				"error": "Unauthorized.",
			})
			ctx.Abort()
			return
		}

		if _, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(initializers.JWTSecret), nil
		}); err != nil {
			ctx.JSON(utils.ISE, gin.H{
				"code":  utils.ISE,
				"error": "Invalid Token.",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
