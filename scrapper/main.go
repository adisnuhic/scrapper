package main

import (
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

	// GRPC SERVER
	go func() {
		listener, err := net.Listen("tcp", cfg.GRPCServerPort)
		if err != nil {
			log.Fatalf("%v", err)
		}
		srv := grpc.NewServer()
		proto.RegisterPostServiceServer(srv, business.MyServer{PostService: postSvc})
		reflection.Register(srv)

		if e := srv.Serve(listener); e != nil {
			log.Fatalf("%v %v", "unable to start GRPC server!", e)
		}
	}()

	// run ever x seconds
	ticker := time.NewTicker(100000 * time.Second)

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
