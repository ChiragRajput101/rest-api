package user

/* This file isolates the database interaction logic
 from rest of the application logic */

import (
	"database/sql"
	"errors"

	"github.com/ChiragRajput101/rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return nil, err
	}

	u := new(types.User) // dynamic allocation of the Type, return *Type (0-val)

	for rows.Next() {
		_, err := scanRowsIntoUserType(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, errors.New("user not found")
	}

	return u, nil
}


func (s *Store) CreateUser(user types.User) error {
	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
	
func scanRowsIntoUserType(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)

	err := rows.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}