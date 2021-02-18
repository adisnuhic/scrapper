package main

import (
	"context"
	"fmt"

	"net"
	"time"

	"github.com/adisnuhic/scrapper/business"
	"github.com/adisnuhic/scrapper/config"
	"github.com/adisnuhic/scrapper/db"
	"github.com/adisnuhic/scrapper/proto"
	"github.com/adisnuhic/scrapper/repositories"
	"github.com/adisnuhic/scrapper/services"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	// Code error code
	Code = "code"
	// Cause error cause
	Cause = "cause"
)

type server struct {
}

func main() {

	// Setup logging
	config.InitLogger()

	// Load config & Init database
	cfg := config.Load()
	db.Init(cfg)

	// New repositories
	postRepo := repositories.NewPostRepository(db.Connection())
	sourceRepo := repositories.NewSourceRepository(db.Connection())

	// New services
	postSvc := services.NewPostService(postRepo)
	sourceSvc := services.NewSourceService(sourceRepo)
	scrapSvc := services.NewScrapService()

	// New businesses
	business.NewPostBusiness(postSvc)
	scrapBl := business.NewScrapBusiness(scrapSvc, postSvc, sourceSvc)

	// Starting GRPC server
	listener, err := net.Listen("tcp", cfg.GRPCServerPort)
	if err != nil {
		log.Fatalf("%v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterPostServiceServer(srv, &server{})
	reflection.Register(srv)
	go func() {
		if e := srv.Serve(listener); e != nil {
			log.Fatalf("%v %v", "unable to start GRPC ", e)
		}
	}()

	// run ever x seconds
	ticker := time.NewTicker(10000 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Printf("%v", "\n ❤ ❤ ❤ Starting scrapper ❤ ❤ ❤ \n")

			// Do scrapping
			_, appErr := scrapBl.Scrap()

			if appErr != nil {
				log.WithFields(log.Fields{
					Code:  appErr.Code,
					Cause: appErr.Cause,
				}).Error(appErr.Message)
			}
		}
	}

}

func (s server) GetAll(ctx context.Context, req *proto.GetAllPostsRequest) (*proto.GetAllPostsResponse, error) {
	mockPosts := []*proto.Post{
		{
			ID:    2,
			Title: "Test1",
			Body:  "Test1 Body",
		},
	}

	return &proto.GetAllPostsResponse{Posts: mockPosts}, nil
}
