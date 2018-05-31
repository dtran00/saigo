package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Customer ...
type Customer struct {
	ID        int       `json:"id" db:"customer_id"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	BirthDate time.Time `json:"birth_date" db:"birth_date"`
	Orders    []*Order
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type dbStore struct {
	db *sqlx.DB
}

type Store interface {
	FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error)
	FindCustomerByID(db *sqlx.DB, id int) (*Customer, error)
	NewCustomer(db *sqlx.DB, email string, first_name string, last_name string, birth_date time.Time) (*Customer, error)
	DeleteCustomer(db *sqlx.DB, id int) error
	UpdateCustomer(db *sqlx.DB, u *Customer) error
	AllCustomers(db *sqlx.DB) ([]*Customer, error)
}

// Refresh ...
func (c *Customer) Refresh(db *sqlx.DB) error {
	return nil
}

// NewCustomer ...
func (s *dbStore) NewCustomer(db *sqlx.DB, email string, first_name string, last_name string, birth_date time.Time) (*Customer, error) {
	_, err := s.db.Exec("INSERT INTO Customers (email, first_name, last_name, birth_date) VALUES ($1, $2, $3, $4)", email, first_name, last_name, birth_date)
	return_customer := Customer{}
	err = s.db.Get(&return_customer, "SELECT * FROM customers WHERE email = $1", email)

	return &return_customer, err
}

// DeleteCustomer ...
func (s *dbStore) DeleteCustomer(db *sqlx.DB, id int) error {
	_, err := s.db.Exec("DELETE FROM Customers WHERE customer_id = $1", id)
	return err
}

// UpdateCustomer ...
func (s *dbStore) UpdateCustomer(db *sqlx.DB, u *Customer) error {
	sqlStatement := `UPDATE Customers SET email = $2, first_name = $3, last_name = $4, birth_date = $5 WHERE customer_id = $1`
	_, err := s.db.Exec(sqlStatement, u.ID, u.Email, u.FirstName, u.LastName, u.BirthDate)
	return err
}

// FindCustomerByEmail ...
func (s *dbStore) FindCustomerByEmail(db *sqlx.DB, email string) (*Customer, error) {
	return_customer := Customer{Email: email}
	err := s.db.Get(&return_customer, "SELECT * FROM Customers WHERE email = $1", email)

	return &return_customer, err
}

// FindCustomerByID ...
func (s *dbStore) FindCustomerByID(db *sqlx.DB, id int) (*Customer, error) {
	return_customer := Customer{ID: id}
	err := s.db.Get(&return_customer, "SELECT * FROM Customers WHERE customer_id = $1", id)

	return &return_customer, err
}

// AllCustomers ...
func (s *dbStore) AllCustomers(db *sqlx.DB) ([]*Customer, error) {
	customers := []*Customer{}

	err := s.db.Select(&customers, "SELECT * FROM Customers")
	return customers, err
}
