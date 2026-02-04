package client

import (
	"net/http"

	crudclient "github.com/mahmud-off/crudUsers/client/internal/crud-client"
	"github.com/mahmud-off/crudUsers/internal/users"
)

type Client interface {
	CreateUser(u *users.User, url string) (*http.Response, error)
	UpdateUser(id int, u *users.User, url string) (*http.Response, error)
	DeleteUser(id int, url string) (*http.Response, error)
	GetUser(id int, url string) (*http.Response, error)
	GETUsers(url string) (*http.Response, error)
}

func NewClient() Client {
	return crudclient.NewCRUDClient()
}
