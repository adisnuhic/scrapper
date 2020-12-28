package controllers

import (
	"github.com/adisnuhic/scrapper/business"
	"github.com/adisnuhic/scrapper/models"
	"github.com/gin-gonic/gin"
)

// IAccountController interface
type IAccountController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
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

	ctrl.RenderSuccess(ctx, nil)
}

// Login user
func (ctrl accountController) Login(ctx *gin.Context) {
	ctrl.RenderSuccess(ctx, "LOGGED IN - NOT IMPLEMENTED")
}
