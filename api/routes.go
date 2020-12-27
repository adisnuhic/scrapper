package main

import "github.com/adisnuhic/scrapper/controllers"

var (
	accountController controllers.IAccountController
)

// initalizeRoutes initialize app routes
func initalizeRoutes() {

	v1 := app.Group("/v1")

	// Auth controller routes
	authRoutes := v1.Group("/account")
	authRoutes.GET("/ping", accountController.Ping)
	authRoutes.POST("/register", accountController.Register)
	authRoutes.POST("/login", accountController.Ping)
}
