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

type User struct {
	Path     string
	Version  string `json:"version" validate:required`
	Username string `json:"username" validate:"required"`
}

func loadUsers(dir string) ([]*User, error) {
	pattern := path.Join(dir, "*")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0, len(files))
	for _, f := range files {
		if path.Base(f) == "README.md" || strings.HasPrefix(path.Base(f), ".") {
			continue
		}

		log.Println("load file", f)

		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		user := new(User)
		if err := yaml.Unmarshal(bytes, user); err != nil {
			return nil, err
		}
		user.Path = f

		users = append(users, user)
	}
	return users, nil
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
