package repository

import (
	"database/sql"
	"errors"
	"final2/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) User {
	var r user
	r.db = db
	return &r
}

func (u *user) Create(newUser entity.User) error {
	db := u.db
	query := `
				INSERT INTO users(fullname, email, password, role, balance, created_at, updated_at)
				VALUES (?, ?, ?, ?, ?, ?, ?)
			`
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	_, err = db.Exec(query, newUser.Fullname, newUser.Email, hashedPwd, "customer", 0, newUser.CreatedAt, newUser.UpdatedAt)
	return err
}

func (u *user) Login(newUser entity.User) error {
	var err error
	db := u.db

	var hashedPwd string
	rows := db.QueryRow("SELECT password FROM users WHERE email = ?", newUser.Email)
	err = rows.Scan(&hashedPwd)
	if err != nil {
		err = errors.New("Incorrect email or password")
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(newUser.Password))

	return err
}

// func (u *user) Update(newUser entity.User) error {
// 	var err error
// 	db := u.db
// 	query := `
// 		UPDATE users SET email = ?,
// 		username = ?,
// 		updated_at = ?
// 		WHERE email = ?
// 	`
// 	_, err = db.Exec(query, newUser.Email, newUser.Username, time.Now())
// 	return err
// }
