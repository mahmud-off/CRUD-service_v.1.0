package database

import (
	"github.com/mahmud-off/crudUsers/internal/services/postgres"
	"github.com/mahmud-off/crudUsers/internal/users"
)

type DB interface {
	DeleteUser(id int) error
	GetUserById(id int) (users.User, error)
	GetUsers() ([]users.User, error)
	AddUser(u *users.User) error
	UpdateUser(u *users.User) error
	Close() error
}

func NewDB(driver string, source string) (DB, error) {
	pq, err := postgres.NewPostgreDB(driver, source)
	return pq, err
}
