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

type Item struct {
	Path        string
	Version     string `json:"version" validate:required`
	ID          string `json:"id" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func loadItems(dir string) ([]*Item, error) {
	files, err := filepath.Glob(path.Join(dir, "*"))
	if err != nil {
		return nil, err
	}

	items := make([]*Item, 0, len(files))
	for _, f := range files {
		if path.Base(f) == "README.md" || strings.HasPrefix(path.Base(f), ".") {
			continue
		}

		log.Println("load file", f)

		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		item := new(Item)
		if err := yaml.Unmarshal(bytes, item); err != nil {
			return nil, err
		}
		item.Path = f

		items = append(items, item)
	}
	return items, nil

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
