package files

import (
	"github.com/go-playground/validator/v10"
)

type Item struct {
	Path        string
	Version     string `json:"version" validate:required`
	ID          string `json:"id" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (i *Item) Validate() error {
	return validator.New().Struct(i)
}

func itemsAsMap(items []*Item) map[string]*Item {
	m := make(map[string]*Item)
	for _, i := range items {
		m[i.ID] = i
	}
	return m
}
