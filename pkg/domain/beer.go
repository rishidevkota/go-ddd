package domain

type Beer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// beer repository
type BeerRepository interface {
	Get(id int64) (*Beer, error)
	GetAll() ([]Beer, error)
	Add(beer *Beer) error
}
