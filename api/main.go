package main

import (
	"github.com/adisnuhic/scrapper/business"
	"github.com/adisnuhic/scrapper/config"
	"github.com/adisnuhic/scrapper/controllers"
	"github.com/adisnuhic/scrapper/db"
	"github.com/adisnuhic/scrapper/initialize"
	"github.com/adisnuhic/scrapper/repositories"
	"github.com/adisnuhic/scrapper/services"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {

	// Load conf
	cfg := config.Load()

	// Init db
	db.Init(cfg)

	// Init repositories
	accountRepo := repositories.NewAccountRepository()

	// Init services
	accountSvc := services.NewAccountService(accountRepo)

	// Init business
	accountBiz := business.NewAccountBusiness(accountSvc)

	// Init controllers
	accountController = controllers.NewAccountController(accountBiz)

	// Init framework
	app = initialize.Gin()

	// Init routes
	initalizeRoutes()

	// Run app
	app.Run(":8282")
}
