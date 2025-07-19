package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password,omitempty"`
}

func (user *User) SaveUser() error {
	query := `
		INSERT INTO users(email, password) 
		VALUES (?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		panic("Prepare Save User failed")
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		panic("Hash Password failed")
	}

	result, err := stmt.Exec(&user.Email, hashedPassword)

	if err != nil {
		panic("Execute Save User failed")
	}

	userID, err := result.LastInsertId()

	user.ID = userID

	return err
}

func GetUsers() ([]User, error) {
	query := "SELECT id, email FROM users"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Email)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retreivePassword string

	err := row.Scan(&u.ID, &retreivePassword)

	if err != nil {
		return errors.New("invalid Credentials")
	}

	isPasswordValid := utils.VerifyPassword(retreivePassword, u.Password)

	if !isPasswordValid {
		return errors.New("invalid Credentials")
	}
	return nil
}
