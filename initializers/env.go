package initializers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	JWTSecret string
	AppURL    string

	host     string
	port     string
	database string
	username string
	password string
	location string
)

func LoadEnv() {
	godotenv.Load()
	AppURL = os.Getenv("APP_URL")

	// database
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	database = os.Getenv("DB_DATABASE")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	location = os.Getenv("DB_LOCATION")

	// get jwt secret
	JWTSecret = os.Getenv("JWT_SECRET")

	// set env mode: production or debug
	gin.SetMode(os.Getenv("GIN_MODE"))
}
