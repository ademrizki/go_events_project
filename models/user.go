package models

import "example.com/rest-api/db"

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
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
	result, err := stmt.Exec(&user.Email, &user.Password)

	if err != nil {
		panic("Execute Save User failed")
	}

	userID, err := result.LastInsertId()

	user.ID = userID

	return err
}