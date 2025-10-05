package initializers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	JWTSecret string
	AppURL    string
)

func LoadEnv() {
	godotenv.Load()
	AppURL = os.Getenv("APP_URL")

	// get jwt secret
	JWTSecret = os.Getenv("JWT_SECRET")

	// set env mode: production or debug
	gin.SetMode(os.Getenv("GIN_MODE"))
}
