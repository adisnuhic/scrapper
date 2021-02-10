package controllers

import (
	"github.com/adisnuhic/scrapper_api/business"
	"github.com/adisnuhic/scrapper_api/viewmodels"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// IPostController intefrace
type IPostController interface {
	GetAll(ctx *gin.Context)
}

type postController struct {
	BaseController
	Business business.IPostBusiness
}

// NewPostController -
func NewPostController(business business.IPostBusiness) IPostController {
	return &postController{
		Business: business,
	}
}

// GetAll returns all posts scrapper service
func (ctrl postController) GetAll(ctx *gin.Context) {

	posts, appErr := ctrl.Business.GetAll()
	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		return
	}

	ctrl.RenderSuccess(ctx, viewmodels.Response{
		Success:   true,
		RequestID: requestid.Get(ctx),
		Data:      posts,
		Error:     nil,
	})

}
