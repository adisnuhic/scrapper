package business

import (
	"context"

	"github.com/adisnuhic/scrapper/proto"
	"github.com/adisnuhic/scrapper/services"
)

type MyServer struct {
	PostService services.PostService
}

// GetAll returns all posts
func (s MyServer) GetAll(ctx context.Context, req *proto.GetAllPostsRequest) (*proto.GetAllPostsResponse, error) {

	dbPosts, _ := s.PostService.GetAll()
	protoPosts := []*proto.Post{}

	for _, post := range *dbPosts {
		protoPosts = append(protoPosts, &proto.Post{ID: post.ID, Title: post.Title, Body: post.Body})
	}

	return &proto.GetAllPostsResponse{Posts: protoPosts}, nil
}
