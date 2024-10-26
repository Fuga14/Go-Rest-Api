package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `bind:"required"`
	Password string `bind:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateLogin() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Login or password is incorrect")
	}

	isValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isValid {
		return errors.New("Login or password is incorrect")
	}

	return nil
}
