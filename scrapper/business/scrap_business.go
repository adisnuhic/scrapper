package business

import (
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
	"github.com/adisnuhic/scrapper/services"
)

// ScrapBusiness interface
type ScrapBusiness interface {
	Scrap() (*models.Posts, *apperror.AppError)
}

type scrapBusiness struct {
	service       services.ScrapService
	sourceService services.SourceService
	postService   services.PostService
}

// NewScrapBusiness new
func NewScrapBusiness(svc services.ScrapService, postSvc services.PostService, sourceSvc services.SourceService) ScrapBusiness {
	return &scrapBusiness{
		service:       svc,
		sourceService: sourceSvc,
		postService:   postSvc,
	}
}

// Scrap the sources
func (bl *scrapBusiness) Scrap() (*models.Posts, *apperror.AppError) {
	// get all sources
	sources, errSources := bl.sourceService.GetAll()
	if errSources != nil {
		return nil, errSources
	}

	// create channel and spin go routines for each source
	c := make(chan *models.Posts)
	for _, src := range *sources {
		go func(src models.Source) {
			posts, _ := bl.service.Scrap(&src)
			c <- posts
		}(src)
	}

	// receive posts from channel
	myPosts := make(models.Posts, 0)
	for range *sources {
		posts := <-c
		if posts != nil {
			for _, postItem := range *posts {
				if postItem.Title != "" {
					myPosts = append(myPosts, postItem)
				}
			}
		}
	}

	// insert collected posts into DB
	dbPosts, errCreate := bl.postService.CreateMany(&myPosts)
	if errCreate != nil {
		return nil, errCreate
	}

	return dbPosts, nil
}
