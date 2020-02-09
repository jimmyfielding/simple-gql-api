package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func New(connection string) (*Db, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func (d *Db) FormatConnection(host, user, dbName string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbName)
}

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (d *Db) GetUserByName(name string) ([]User, error) {
	query, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		return nil, fmt.Errorf("GetUserByName returned an error: ", err)
	}

	rows, err := query.Query(name)
	if err != nil {
		return nil, fmt.Errorf("GetUserByName returned an error: ", err)
	}

	var u User
	users := []User{}
	for rows.Next() {
		err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Age,
			&u.Profession,
			&u.Friendly,
		)
		if err != nil {
			return nil, fmt.Errorf("GetUserByName returns an error: ", err)
		}

		users = append(users, u)
	}

	return users, nil
}
