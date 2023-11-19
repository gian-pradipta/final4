package repository

import (
	"database/sql"
	"final2/internal/entity"
)

type transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	var repo transaction
	repo.db = db
	return &repo
}

func (t *transaction) validateTransaction(transaction entity.Transaction) bool {
	var valid bool = true
	db := t.db
	prepo := NewProduct(db)
	urepo := NewUserRepository(db)
	product, err := prepo.Get(transaction.ProductId)
	if err != nil {
		valid = false
		return valid
	}
	user, err := urepo.Get(transaction.UserId)
	if err != nil {
		valid = false
		return valid
	}

	if product.Stock < transaction.Quantity {
		return false
	}
	if product.Price*transaction.Quantity > user.Balance {
		return false
	}
	return valid
}

// func (t *transaction) Create(transaction entity.Transaction) error {
// 	var err error
// 	db := t.db
// 	valid := t.validateTransaction(transaction)
// 	prepo := NewProduct(db)
// 	if !valid {
// 		return err
// 	}

// 	_, err = db.Exec(`INSERT INTO transaction_history (product_id, user_id, quantity, total_price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`, transaction.ProductId, transaction.UserId, transaction.Quantity, transaction.TotalPrice, time.Now(), time.Now())
// 	product, err := prepo.Get(transaction.ProductId)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = db.Exec(`UPDATE category SET sold_product_amount = ? WHERE id = ?`, product.CategoryId)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

func (t *transaction) Create(transaction entity.Transaction) error {
	var err error
	db := t.db
	_, err = db.Exec(`INSERT INTO transaction_history (product_id, user_id, quantity, total_price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`, transaction.ProductId, transaction.UserId, transaction.Quantity, transaction.TotalPrice, transaction.CreatedAt, transaction.UpdatedAt)
	if err != nil {
		return err
	}
	prepo := NewProduct(db)
	product, err := prepo.Get(transaction.ProductId)
	if err != nil {
		return err
	}
	_, err = db.Exec(`UPDATE category SET sold_product_amount = sold_product_amount + ? WHERE id = ?`, transaction.Quantity, product.CategoryId)
	if err != nil {
		return err
	}
	_, err = db.Exec(`UPDATE users SET balance = balance - ? WHERE id = ?`, transaction.TotalPrice, transaction.UserId)
	return err
}

func (t *transaction) GetUserProduct(userEmail string, productId int) (entity.User, entity.Product, error) {
	var err error
	var user entity.User
	var product entity.Product
	var userId int

	db := t.db
	prepo := NewProduct(db)
	urepo := NewUserRepository(db)

	product, err = prepo.Get(productId)
	rows := db.QueryRow(`SELECT id FROM users WHERE email = ?`, userEmail)
	err = rows.Scan(&userId)
	if err != nil {
		return user, product, err
	}
	user, err = urepo.Get(userId)
	return user, product, err
}
