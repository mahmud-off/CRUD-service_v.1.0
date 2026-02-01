package crud_server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mahmud-off/crudUsers/internal/users"
)

func (s *CRUDServer) UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			if ok := r.Header.Get("id"); len(ok) == 0 {
				users, err := s.DB.GetUsers()
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				var out string
				for _, user := range users {
					personalInfo := fmt.Sprintf("ID : %d | Name : %s | Email : %s;\n", user.ID, user.Name, user.Email)
					out += personalInfo
				}
				w.Write([]byte(out))
			} else {
				id, err := strconv.Atoi(ok)
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				user, err := s.DB.GetUserById(id)
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				var out string
				personalInfo := fmt.Sprintf("ID : %d | Name : %s | Email : %s;\n", user.ID, user.Name, user.Email)
				out += personalInfo
				w.Write([]byte(out))
			}

		}
	case http.MethodPost:
		{
			u := &users.User{
				Name:     r.Header.Get("name"),
				Email:    r.Header.Get("email"),
				Password: r.Header.Get("password"),
			}
			err := s.DB.AddUser(u)
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	case http.MethodPatch:
		{
			if ok := r.Header.Get("id"); len(ok) != 0 {
				id, err := strconv.Atoi(ok)
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				u := &users.User{
					ID:       id,
					Name:     r.Header.Get("name"),
					Email:    r.Header.Get("email"),
					Password: r.Header.Get("password"),
				}

				err = s.DB.UpdateUser(u)
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		}
	case http.MethodDelete:
		{
			if ok := r.Header.Get("id"); len(ok) != 0 {
				id, err := strconv.Atoi(ok)
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				err = s.DB.DeleteUser(id)
				if err != nil {
					log.Fatal(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	default:
		{
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Hello, World!"))
		}
	}
}
