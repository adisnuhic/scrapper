package controllers

import (
	"github.com/adisnuhic/scrapper/business"
	"github.com/adisnuhic/scrapper/models"
	"github.com/gin-gonic/gin"
)

// IAccountController interface
type IAccountController interface {
	Ping(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type accountController struct {
	BaseController
	Business business.IAccountBusiness
}

// NewAccountController -
func NewAccountController(business business.IAccountBusiness) IAccountController {
	return &accountController{
		Business: business,
	}
}

// Ping returns string "pong"
func (ctrl accountController) Ping(ctx *gin.Context) {
	msg := ctrl.Business.Ping()
	ctrl.RenderSuccess(ctx, msg)
}

// Register user
func (ctrl accountController) Register(ctx *gin.Context) {
	reqObj := &models.User{}

	if err := ctx.ShouldBindJSON(reqObj); err != nil {
		ctrl.RednerBadRequest(ctx, err.Error())
		return
	}

	appErr := ctrl.Business.Register(reqObj.FirstName, reqObj.LastName, reqObj.Email, reqObj.Password)
	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		return
	}

	ctrl.RenderSuccess(ctx, "SUCCESS")
}
