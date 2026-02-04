package crudclient

import "net/http"

type CRUDClient struct {
	client *http.Client
}

func NewCRUDClient() *CRUDClient {
	return &CRUDClient{client: &http.Client{}}
}
