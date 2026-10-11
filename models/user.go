package models

import (
	"errors"

	"example.com/go-rest-api/db"
	"example.com/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) Values (?, ?)`
	stmt, err := db.DB.Prepare(query) // complex query, so we prepare it first to avoid SQL injection and improve performance
	if err != nil {
		return err
	}
	defer stmt.Close()                                    // close the statement after execution
	hashedPassword, err := utils.HashPassword(u.Password) // hash the password before saving it to the database
	result, err := stmt.Exec(u.Email, hashedPassword)     // execute the statement with the user data
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users where email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	row.Scan(&u.ID, &retrievedPassword)
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
