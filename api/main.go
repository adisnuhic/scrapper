package main

import (
	"github.com/adisnuhic/scrapper_api/business"
	"github.com/adisnuhic/scrapper_api/config"
	"github.com/adisnuhic/scrapper_api/controllers"
	"github.com/adisnuhic/scrapper_api/db"
	"github.com/adisnuhic/scrapper_api/initialize"
	"github.com/adisnuhic/scrapper_api/repositories"
	"github.com/adisnuhic/scrapper_api/services"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {

	// Load conf
	cfg := config.Load()

	// Init db
	db.Init(cfg)

	// Init repositories
	accountRepo := repositories.NewAccountRepository(db.Connection())
	userRepo := repositories.NewUserRepository(db.Connection())
	authRepo := repositories.NewAuthProviderRepository(db.Connection())
	tokenRepo := repositories.NewTokenRepository(db.Connection())

	// Init services
	accountSvc := services.NewAccountService(accountRepo)
	userSvc := services.NewUserService(userRepo)
	authProviderSvc := services.NewAuthProviderService(authRepo)
	authSvc := services.NewAuthService()
	tokenSv := services.NewTokenService(tokenRepo)
	postSvc := services.NewPostService()

	// Init business
	accountBiz := business.NewAccountBusiness(accountSvc, userSvc, authProviderSvc, authSvc, tokenSv)
	business.NewUserBusiness(userSvc)
	postBiz := business.NewPostBusiness(postSvc)

	// Init controllers
	accountController = controllers.NewAccountController(accountBiz)
	postController = controllers.NewPostController(postBiz)

	// Init framework
	app = initialize.Gin()

	// Init routes
	initalizeRoutes()

	// Run app
	app.Run(":8282")
}
