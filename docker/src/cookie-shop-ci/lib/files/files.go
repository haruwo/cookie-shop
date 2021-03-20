package files

import (
	"fmt"
	"path"
	"strings"
)

type Files struct {
	Users []*User
	Items []*Item
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