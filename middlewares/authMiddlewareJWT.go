package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/what-crud/controllers"
	"github.com/what-crud/initializers"
	"github.com/what-crud/utils"
)

func AuthMiddlewareJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("Authorization")

		tokenString := strings.Split(tokenHeader, "Bearer ")[1]
		// tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")
		// tokenString = strings.TrimSpace(tokenString)

		if tokenString == "" {
			ctx.JSON(utils.UNAUTH, gin.H{
				"code":  utils.UNAUTH,
				"error": "Unauthorized.",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(initializers.JWTSecret), nil
		})
		if err != nil {
			ctx.JSON(utils.ISE, gin.H{
				"code":    utils.ISE,
				"message": "Invalid Token.",
				"error":   err.Error(),
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(utils.UNAUTH, gin.H{
				"code":  utils.UNAUTH,
				"error": "Unauthorized.",
			})
			ctx.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(utils.UNAUTH, gin.H{
				"code":  utils.UNAUTH,
				"error": "Expired Token!",
			})
			ctx.Abort()
			return
		}

		user, err := controllers.GetUserByAuthID(claims["userId"].(string), ctx)
		if err != nil {
			ctx.JSON(utils.NF, gin.H{
				"code":    utils.NF,
				"message": "User Not Found!",
				"error":   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("authUser", user)

		ctx.Next()
	}
}
