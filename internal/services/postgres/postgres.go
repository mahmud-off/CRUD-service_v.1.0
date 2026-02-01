package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/mahmud-off/crudUsers/internal/users"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgreDB(driver string, source string) (*Postgres, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) DeleteUser(id int) error {
	_, err := p.db.Exec("delete from users where id=$1", id)
	return err
}

func (p *Postgres) GetUserById(id int) (users.User, error) {
	var u users.User = users.User{}
	err := p.db.QueryRow("select * from users where id=$1", id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.TimeStamp)
	return u, err
}

func (p *Postgres) GetUsers() ([]users.User, error) {

	rows, err := p.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u []users.User = make([]users.User, 0)
	for rows.Next() {
		var unit users.User = users.User{}
		err := rows.Scan(&unit.ID, &unit.Name, &unit.Email, &unit.Password, &unit.TimeStamp)
		if err != nil {
			return nil, err
		}
		u = append(u, unit)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return u, nil

}

func (p *Postgres) AddUser(u *users.User) error {
	_, err := p.db.Exec("insert into users (name, email, password) values ($1, $2, $3)", u.Name, u.Email, u.Password)
	return err
}

func (p *Postgres) UpdateUser(u *users.User) error {
	_, err := p.db.Exec("update users SET name=$2, email=$3, password=$4 where id=$1;", u.ID, u.Name, u.Email, u.Password)
	return err
}
