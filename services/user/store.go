package user

/* This file isolates the database interaction logic
from rest of the application logic */

import (
	"database/sql"
	"fmt"
	"log"

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

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?,?,?,?)", 
											user.FirstName, user.LastName, user.Email, user.Password)
	
	if err != nil {
		return nil
	}			
	return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		log.Fatal("query err")
		return nil, err
	}

	// fmt.Println("query ok")

	u := new(types.User) // dynamic allocation of the Type, return *Type (0-val)

	for rows.Next() {
		user, err := scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
		u = user
	}

	// fmt.Println("ok1")

	if u.ID == 0 || u == nil {
		return nil, fmt.Errorf("not found")
	}

	// fmt.Println("ok2")

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	u := new(types.User) // dynamic allocation of the Type, return *Type (0-val)

	for rows.Next() {
		user, err := scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
		u = user
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("not found")
	}

	return u, nil
}
	
func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}