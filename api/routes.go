package main

import "github.com/adisnuhic/scrapper/controllers"

var (
	accountController controllers.IAccountController
)

// initalizeRoutes initialize app routes
func initalizeRoutes() {

	v1 := app.Group("/v1")

	// Auth controller routes
	accountRoutes := v1.Group("/account")
	accountRoutes.POST("/register", accountController.Register)
	accountRoutes.POST("/login", accountController.Login)
}
