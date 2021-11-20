package main

import (
	"fmt"

	"github.com/rishidevkota/go-ddd/pkg/application"
	"github.com/rishidevkota/go-ddd/pkg/domain"
	"github.com/rishidevkota/go-ddd/pkg/infrastructure/presistence"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Beer{})

	var beerStore domain.BeerRepository = presistence.NewBeerStorage(db)
	service := application.NewBeerService(beerStore)

	b := domain.Beer{Name: "beer1"}
	service.AddBeer(&b)

	c, _ := service.GetBeer(1)
	fmt.Println(c, b)

}
