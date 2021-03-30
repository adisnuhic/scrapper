package controllers

import (
	"github.com/adisnuhic/scrapper_api/business"
	"github.com/adisnuhic/scrapper_api/proto"
	"github.com/adisnuhic/scrapper_api/viewmodels"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// IPostController intefrace
type IPostController interface {
	GetAll(ctx *gin.Context)
}

type postController struct {
	BaseController
	Business business.IPostBusiness
	GrpcConn grpc.ClientConnInterface
}

// NewPostController -
func NewPostController(business business.IPostBusiness, grpcConn grpc.ClientConnInterface) IPostController {
	return &postController{
		Business: business,
		GrpcConn: grpcConn,
	}
}

// GetAll returns all posts scrapper service
func (ctrl postController) GetAll(ctx *gin.Context) {

	// grpc client call
	req := &proto.GetAllPostsRequest{}
	client := proto.NewPostServiceClient(ctrl.GrpcConn)

	posts, appErr := client.GetAll(ctx, req)
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
