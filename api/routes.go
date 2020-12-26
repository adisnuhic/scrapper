package main

import "github.com/adisnuhic/scrapper/controllers"

var (
	pingController controllers.IPingController
)

// initalizeRoutes initialize app routes
func initalizeRoutes() {

	v1 := app.Group("/v1")

	// Ping controller routes
	pingRoutes := v1.Group("/ping")
	pingRoutes.GET("/", pingController.Ping)
}
