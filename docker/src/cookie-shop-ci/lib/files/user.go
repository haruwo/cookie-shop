package files

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Path     string
	Version  string `json:"version" validate:required`
	Username string `json:"username" validate:"required"`
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}

func usersAsMap(users []*User) map[string]*User {
	m := make(map[string]*User)
	for _, u := range users {
		m[u.Username] = u
	}
	return m
}
