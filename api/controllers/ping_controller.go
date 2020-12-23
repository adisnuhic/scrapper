package controllers

import (
	"github.com/adisnuhic/scrapper/business"
	"github.com/gin-gonic/gin"
)

// IPingController interface
type IPingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct {
	Business business.IPingBusiness
}

// NewPingController -
func NewPingController(business business.IPingBusiness) IPingController {
	return &pingController{
		Business: business,
	}
}

// Ping returns string "pong"
func (ctrl pingController) Ping(ctx *gin.Context) {
	msg := ctrl.Business.Ping()

	ctx.JSON(200, gin.H{
		"message": msg,
	})
}
