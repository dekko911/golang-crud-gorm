package routes

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/what-crud/utils"
)

func keyFunc(ctx *gin.Context) string {
	return ctx.ClientIP()
}

func errorHandler(ctx *gin.Context, info ratelimit.Info) {
	ctx.JSON(utils.TMR, gin.H{
		"code":  utils.TMR,
		"error": "Too many requests. Try again in " + time.Until(info.ResetTime).String(),
	})
}

func RateLimit(rate time.Duration, limit uint) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  rate,
		Limit: limit,
	})

	return ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
}
