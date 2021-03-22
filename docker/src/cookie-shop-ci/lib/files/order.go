package files

import (
	"github.com/go-playground/validator/v10"
)

type Order struct {
	Path     string
	Version  string       `json:"version" validate:required`
	Username string       `json:"username" validate:"required"`
	Items    []*OrderItem `json:"items" validate:"required,dive"`
}

type OrderItem struct {
	Path   string
	ID     string `json:"id" validate:"required"`
	Amount int64  `json:"amount" validate:"required,gte=1,lte=96"`
}

func (o *Order) Validate() error {
	return validator.New().Struct(o)
}
