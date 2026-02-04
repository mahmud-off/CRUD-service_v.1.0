package crudclient

import (
	"net/http"
	"strconv"

	"github.com/mahmud-off/crudUsers/internal/users"
)

func (c *CRUDClient) CreateUser(u *users.User, url string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("name", u.Name)
	req.Header.Add("email", u.Email)
	req.Header.Add("password", u.Password)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CRUDClient) UpdateUser(id int, u *users.User, url string) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("id", strconv.Itoa(id))
	req.Header.Add("name", u.Name)
	req.Header.Add("email", u.Email)
	req.Header.Add("password", u.Password)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CRUDClient) DeleteUser(id int, url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("id", strconv.Itoa(id))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CRUDClient) GetUser(id int, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("id", strconv.Itoa(id))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CRUDClient) GETUsers(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
