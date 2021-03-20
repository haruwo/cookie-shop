package files

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type Order struct {
	Path     string
	Version  string       `json:"version" validate:required`
	Username string       `json:"username" validate:"required"`
	Items    []*OrderItem `json:"items" validate:"required"`
}

type OrderItem struct {
	Path   string
	ID     string `json:"id" validate:"required"`
	Amount int64  `json:"amount" validate:"required"`
}

func loadOrders(dir string) ([]*Order, error) {
	files, err := filepath.Glob(path.Join(dir, "*"))
	if err != nil {
		return nil, err
	}

	orders := make([]*Order, 0, len(files))
	for _, f := range files {
		if path.Base(f) == "README.md" || strings.HasPrefix(path.Base(f), ".") {
			continue
		}

		log.Println("load file", f)

		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		order := new(Order)
		if err := yaml.Unmarshal(bytes, order); err != nil {
			return nil, err
		}
		order.Path = f

		orders = append(orders, order)
	}
	return orders, nil
}

func (o *Order) Validate() error {
	return validator.New().Struct(o)
}
