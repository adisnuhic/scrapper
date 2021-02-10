package main

import "github.com/adisnuhic/scrapper_api/controllers"

var (
	accountController controllers.IAccountController
	postController    controllers.IPostController
)

// initalizeRoutes initialize app routes
func initalizeRoutes() {

	v1 := app.Group("/v1")

	// Auth controller routes
	accountRoutes := v1.Group("/account")
	accountRoutes.POST("/register", accountController.Register)
	accountRoutes.POST("/login", accountController.Login)
	accountRoutes.POST("/refresh-token", accountController.RefreshToken)

	// Post controller routes
	postRoutes := v1.Group("/posts")
	postRoutes.GET("/", postController.GetAll)

}
