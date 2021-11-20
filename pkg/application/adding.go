package application

import (
	"github.com/rishidevkota/go-ddd/pkg/domain"
)

// beer service
type BeerService interface {
	AddBeer(beer *domain.Beer) error
	GetBeer(id int64) (*domain.Beer, error)
}

type beerService struct {
	bR domain.BeerRepository
}

// NewBeerService returns a new beer service
func NewBeerService(bR domain.BeerRepository) BeerService {
	return &beerService{bR}
}

func (s *beerService) AddBeer(beer *domain.Beer) error {
	return s.bR.Add(beer)
}

func (s *beerService) GetBeer(id int64) (*domain.Beer, error) {
	return s.bR.Get(id)
}
