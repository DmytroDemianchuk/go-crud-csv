package services

import "github.com/dmytrodemianchuk/go-crud-csv/internal/repository"

type Services struct {
	Products *ServiceProducts
}

func New(repo *repository.Repository) *Services {
	return &Services{Products: initProducts(repo.Products)}
}
