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

func (u *user) Login(newUser entity.User) (string, error) {
	var err error
	var group string
	db := u.db

	var hashedPwd string
	rows := db.QueryRow("SELECT password, role FROM users WHERE email = ?", newUser.Email)
	err = rows.Scan(&hashedPwd, &group)
	if err != nil {
		err = errors.New("Incorrect email or password")
		return group, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(newUser.Password))

	return group, err
}

func (u *user) TopUp(user entity.User) error {
	var err error
	db := u.db

	_, err = db.Exec("UPDATE users SET balance = ? WHERE email = ?", user.Balance, user.Email)
	return err
}
