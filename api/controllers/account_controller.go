package controllers

import (
	"github.com/adisnuhic/scrapper_api/business"
	"github.com/adisnuhic/scrapper_api/viewmodels"
	"github.com/gin-gonic/gin"
)

// IAccountController interface
type IAccountController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
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
	type ReqObj struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}
	reqObj := &ReqObj{}

	if err := ctx.ShouldBindJSON(reqObj); err != nil {
		ctrl.RednerBadRequest(ctx, err.Error())
		return
	}

	accessToken, refreshToken, appErr := ctrl.Business.Register(reqObj.FirstName, reqObj.LastName, reqObj.Email, reqObj.Password)
	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		return
	}

	ctrl.RenderSuccess(ctx, viewmodels.Auth{
		User:         nil,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}

// Login user
func (ctrl accountController) Login(ctx *gin.Context) {

	type ReqLogin struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var reqObj ReqLogin

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		ctrl.RednerBadRequest(ctx, err.Error())
		return
	}

	user, accessToken, refreshToken, appErr := ctrl.Business.Login(reqObj.Email, reqObj.Password)

	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		return
	}

	ctrl.RenderSuccess(ctx, &viewmodels.Auth{
		User:         user,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}

// RefreshToken -
func (ctrl accountController) RefreshToken(ctx *gin.Context) {
	type ReqRefreshToken struct {
		Token string `json:"token" binding:"required"`
	}

	var reqObj ReqRefreshToken

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		ctrl.RednerBadRequest(ctx, err.Error())
		return
	}

	user, accessToken, refreshToken, appErr := ctrl.Business.RefreshToken(reqObj.Token)
	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		return
	}

	ctrl.RenderSuccess(ctx, &viewmodels.Auth{
		User:         user,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})

}
