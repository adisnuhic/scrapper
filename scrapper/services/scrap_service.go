package services

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/adisnuhic/scrapper/ecode"
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
)

const (
	// SourceAutoBlog source
	SourceAutoBlog = "autoblog"
	// SourceBuzzFeed source
	SourceBuzzFeed = "buzzfeed"
)

// ScrapService -
type ScrapService interface {
	Scrap(source *models.Source) (*models.Posts, *apperror.AppError)
}

type scrapService struct{}

// NewScrapService -
func NewScrapService() ScrapService {
	return &scrapService{}
}

// Scrap the sources
func (svc *scrapService) Scrap(source *models.Source) (*models.Posts, *apperror.AppError) {
	// client
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// request
	request, errRequest := http.NewRequest("GET", source.SourceURL, nil)
	if errRequest != nil {
		return nil, apperror.New(ecode.ErrUnableToMakeNewRequestCode, errRequest, ecode.ErrUnableToMakeNewRequestMsg)
	}

	// response
	response, errResponse := client.Do(request)
	if errResponse != nil {
		return nil, apperror.New(ecode.ErrUnableToGetFromSourceCode, errResponse, ecode.ErrUnableToGetFromSourceMsg)
	}

	if response.StatusCode != 200 {
		return nil, apperror.New(ecode.ErrUnableToGetFromSourceCode, errResponse, ecode.ErrUnableToGetFromSourceMsg)
	}

	// close body at the end of func
	defer response.Body.Close()

	// make a goQuery document
	doc, errDoc := goquery.NewDocumentFromReader(response.Body)
	if errDoc != nil {
		return nil, apperror.New(ecode.ErrUnableToReadSourceBodyCode, errDoc, ecode.ErrUnableToReadSourceBodyMsg)
	}

	// AutoBlog
	if source.Source == SourceAutoBlog {
		posts := extractData(".record_details", ".record-heading span", ".subTitle", doc)
		return posts, nil
	}

	// BuzzFeed
	if source.Source == SourceBuzzFeed {
		posts := extractData(".story-card", ".js-card__link", ".js-card__description", doc)
		return posts, nil
	}

	return &models.Posts{}, nil
}

// Extract data from document
func extractData(rootElement string, titleElement string, bodyElement string, doc *goquery.Document) *models.Posts {
	posts := make(models.Posts, 0)

	doc.Find(rootElement).Each(func(i int, s *goquery.Selection) {
		// for each item found get title and body
		title := s.Find(titleElement).Text()
		body := s.Find(bodyElement).Text()

		p := &models.Post{}
		p.Title = title
		p.Body = body

		posts = append(posts, *p)
	})

	return &posts
}
