package domain

type Review struct {
	ID     int64 `json:"id"`
	BeerID int64 `json:"beer_id"`
	Score  int64 `json:"score"`
}

// review repository
type ReviewRepository interface {
	// get all bear reviews
	GetAll(id int64) ([]*Review, error)
	Create(review *Review) error
}
