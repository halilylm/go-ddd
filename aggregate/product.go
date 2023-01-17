package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/halilylm/go-ddd/entity"
)

type Product struct {
	item     *entity.Item
	price    int
	quantity int
}

var (
	ErrMissingValues = errors.New("missing values to create a product")
)

func NewProduct(name, description string, price int) (*Product, error) {
	if name == "" || description == "" {
		return nil, ErrMissingValues
	}
	return &Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p *Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p *Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) GetPrice() int {
	return p.price
}