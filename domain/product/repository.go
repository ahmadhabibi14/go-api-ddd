package product

import (
	"errors"
	"go-api-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound 			= errors.New("the product was not found")
	ErrProductAlreadyExist	= errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}