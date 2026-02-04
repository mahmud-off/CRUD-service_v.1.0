package main

import (
	"fmt"
	"io"
	"log"

	"github.com/mahmud-off/crudUsers/client/internal/client"
)

func main() {
	client := client.NewClient()

	// u := &users.User{
	// 	Name:     "vadikVad",
	// 	Email:    "pupupu@dog.com",
	// 	Password: "123vadimkapupkin321",
	// }

	//resp, err := client.CreateUser(u, "http://localhost:8080/user")
	resp, err := client.GETUsers("http://localhost:8080/user")
	//resp, err := client.UpdateUser(3, &users.User{}, "http://localhost:8080/user")
	//resp, err := client.DeleteUser(3, "http://localhost:8080/user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(text))
}
