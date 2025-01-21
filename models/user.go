package models

import (
	"errors"
	"log"

	"github.com/dsabljic/event-management/db"
	"github.com/dsabljic/event-management/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, u.Email, hashedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	log.Printf("User saved with ID: %d\n", u.ID)
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = $1`
	var retrievedPassword string

	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
