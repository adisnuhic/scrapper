package main

import (
	"github.com/adisnuhic/scrapper/business"
	"github.com/adisnuhic/scrapper/config"
	"github.com/adisnuhic/scrapper/controllers"
	"github.com/adisnuhic/scrapper/db"
	"github.com/adisnuhic/scrapper/repositories"
	"github.com/adisnuhic/scrapper/services"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load conf
	cfg := config.Load()

	// Init db
	db.Init(cfg)

	// Init repositories
	pingRepo := repositories.NewPingRepository()

	// Init services
	pingService := services.NewPingService(pingRepo)

	// Init business
	blPing := business.NewPingBusiness(pingService)

	// Init controllers
	pingController := controllers.NewPingController(blPing)

	// Init framework
	app := gin.Default()

	/* ------------------------------------------------- */
	/*   					ROUTES						 */
	/* ------------------------------------------------- */
	// Ping controller routes
	pingRoutes := app.Group("/ping")
	pingRoutes.GET("/", pingController.Ping)

	// Run app
	app.Run(":8282")
}
