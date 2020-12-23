package business

import "github.com/adisnuhic/scrapper/services"

// IPingBusiness interface
type IPingBusiness interface {
	Ping() string
}

type pingBusiness struct {
	Service services.IPingService
}

// NewPingBusiness -
func NewPingBusiness(svc services.IPingService) IPingBusiness {
	return &pingBusiness{
		Service: svc,
	}
}

// Ping returns string "pong"
func (bl pingBusiness) Ping() string {
	return bl.Service.Ping()
}
