package main

import (
	"github.com/what-crud/initializers"
	"github.com/what-crud/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	// routes
	r := routes.ApiRoutes()

	r.Run()
}
