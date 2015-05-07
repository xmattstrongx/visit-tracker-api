package main

import (
	"RestApiProject/controllers"
	"RestApiProject/logger"

	"net/http"
)

// GET http://localhost:8080/secret
// and use admin,admin for the credentials

func main() {

	controllers.RegisterVisitorApiRoutes()
	logger.Message("Started server")
	http.ListenAndServe(":8080", nil)
}
