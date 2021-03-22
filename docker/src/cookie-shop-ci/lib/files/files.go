package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Files struct {
	Users  []*User
	Items  []*Item
	Orders []*Order
}

func Open(dir string) (*Files, error) {
	users, err := loadUsers(path.Join(dir, "users"))
	if err != nil {
		return nil, err
	}

	items, err := loadItems(path.Join(dir, "items"))
	if err != nil {
		return nil, err
	}

	orders, err := loadOrders(path.Join(dir, "orders"))
	if err != nil {
		return nil, err
	}

	return &Files{
		Users:  users,
		Items:  items,
		Orders: orders,
	}, nil
}

func (f *Files) Validate() error {
	errors := make([]error, 0)
	for _, u := range f.Users {
		if err := u.Validate(); err != nil {
			errors = append(errors, fmt.Errorf("[%s] %w", u.Path, err))
		}
	}

	for _, i := range f.Items {
		if err := i.Validate(); err != nil {
			errors = append(errors, fmt.Errorf("[%s] %w", i.Path, err))
		}
	}

	for _, o := range f.Orders {
		if err := o.Validate(); err != nil {
			errors = append(errors, fmt.Errorf("[%s] %w", o.Path, err))
		}
	}

	usersMap := usersAsMap(f.Users)
	itemsMap := itemsAsMap(f.Items)

	for _, order := range f.Orders {
		if _, found := usersMap[order.Username]; !found {
			errors = append(errors, fmt.Errorf("[%s] not found user %s", order.Path, order.Username))
		}

		for _, oi := range order.Items {
			if _, found := itemsMap[oi.ID]; !found {
				errors = append(errors, fmt.Errorf("[%s] not found item %s", order.Path, oi.ID))
			}
		}
	}

	if len(errors) > 0 {
		msgs := make([]string, 0, len(errors))
		for _, err := range errors {
			msgs = append(msgs, err.Error())
		}
		return fmt.Errorf("found validation errors:\n" + strings.Join(msgs, "\n"))
	}

	return nil
}

func (f *Files) FindOrder(id string) *Order {
	for _, o := range f.Orders {
		if path.Base(o.Path) == id {
			return o
		}
	}
	return nil
}

func (f *Files) FindItem(id string) *Item {
	for _, i := range f.Items {
		if i.ID == id {
			return i
		}
	}
	return nil
}

func loadUsers(dir string) ([]*User, error) {
	var users []*User

	err := loadFiles(dir, loadFilesHandler{
		EnsureCap: func(cap int) {
			users = make([]*User, 0, cap)
		},

		LoadFile: func(f string, data []byte) error {
			user := new(User)
			user.Path = f
			err := yaml.Unmarshal(data, user)
			if err == nil {
				users = append(users, user)
			}
			return err
		},
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func loadItems(dir string) ([]*Item, error) {
	var items []*Item

	err := loadFiles(dir, loadFilesHandler{
		EnsureCap: func(cap int) {
			items = make([]*Item, 0, cap)
		},

		LoadFile: func(f string, data []byte) error {
			item := new(Item)
			item.Path = f
			err := yaml.Unmarshal(data, item)
			if err == nil {
				items = append(items, item)
			}
			return err
		},
	})

	if err != nil {
		return nil, err
	}

	return items, nil
}

func loadOrders(dir string) ([]*Order, error) {
	var orders []*Order

	err := loadFiles(dir, loadFilesHandler{
		EnsureCap: func(cap int) {
			orders = make([]*Order, 0, cap)
		},

		LoadFile: func(f string, data []byte) error {
			order := new(Order)
			order.Path = f
			err := yaml.Unmarshal(data, order)
			if err == nil {
				orders = append(orders, order)
			}
			return err
		},
	})

	if err != nil {
		return nil, err
	}

	return orders, nil
}

type loadFilesHandler struct {
	EnsureCap func(cap int)
	LoadFile  func(path string, data []byte) error
}

func loadFiles(dir string, handler loadFilesHandler) error {
	files, err := filepath.Glob(path.Join(dir, "*"))
	if err != nil {
		return err
	}

	handler.EnsureCap(len(files))

	for _, f := range files {
		if path.Base(f) == "README.md" || strings.HasPrefix(path.Base(f), ".") {
			continue
		}

		log.Println("load file", f)

		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		if err := handler.LoadFile(f, bytes); err != nil {
			return err
		}
	}
	return nil
}
