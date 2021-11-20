package presistence

import (
	"github.com/rishidevkota/go-ddd/pkg/domain"
	"gorm.io/gorm"
)

// Beer Storage
type BeerStorage struct {
	db *gorm.DB
}

// NewBeerStorage creates a new BeerStorage
func NewBeerStorage(db *gorm.DB) *BeerStorage {
	return &BeerStorage{db: db}
}

// Add domain.Beer to the storage
func (s *BeerStorage) Add(beer *domain.Beer) error {
	return s.db.Create(beer).Error
}

// Get domain.Beer from the storage
func (s *BeerStorage) Get(id int64) (*domain.Beer, error) {
	var beer domain.Beer
	if err := s.db.First(&beer, id).Error; err != nil {
		return nil, err
	}
	return &beer, nil
}

// GetAll domain.Beer from the storage
func (s *BeerStorage) GetAll() ([]domain.Beer, error) {
	var beers []domain.Beer
	if err := s.db.Find(&beers).Error; err != nil {
		return nil, err
	}
	return beers, nil
}
