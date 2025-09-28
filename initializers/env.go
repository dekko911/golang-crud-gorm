package initializers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	godotenv.Load()

	// set env mode: production or debug
	gin.SetMode(os.Getenv("GIN_MODE"))
}
