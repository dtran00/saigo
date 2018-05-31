package models

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite

	store  *dbStore
	db     *sqlx.DB
	assert *assert.Assertions
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=models sslmode=disable"
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.store.db.Query("DELETE FROM Customers")
	if err != nil {
		s.T().Fatal(err)
	}
	s.assert = assert.New(s.T())
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestNewCustomer() {
	customer, err := s.store.NewCustomer(s.store.db, "email", "first", "last", time.Now())
	s.assert.Nil(err)
	s.assert.NotNil(customer)
	s.assert.Equal(customer.Email, "email")
}

func (s *StoreSuite) TestDeleteCustomer() {
	_, err := s.store.db.Exec(`INSERT INTO customers (email, first_name, last_name, birth_date) VALUES ('test','test','test', NOW());`)
	s.assert.Nil(err)

	var id int
	err = s.store.db.QueryRow(`SELECT customer_id FROM Customers WHERE email='test';`).
		Scan(&id)

	err = s.store.DeleteCustomer(s.store.db, id)
	s.assert.Nil(err)
}

func (s *StoreSuite) TestUpdateCustomer() {
	_, err := s.store.db.Exec(`INSERT INTO customers (email, first_name, last_name, birth_date) VALUES ('test','test','test', NOW());`)
	s.assert.Nil(err)

	var id int
	err = s.store.db.QueryRow(`SELECT customer_id FROM Customers WHERE email='test';`).
		Scan(&id)

	customer, err := s.store.FindCustomerByID(s.store.db, id)
	s.assert.Nil(err)

	customer.Email = "changed_email"

	err = s.store.UpdateCustomer(s.store.db, customer)
	s.assert.Nil(err)
	s.assert.Equal(customer.Email, "changed_email")
}

func (s *StoreSuite) TestFindCustomerByEmail() {
	_, err := s.store.db.Exec(`INSERT INTO customers (email, first_name, last_name, birth_date) VALUES ('test_email','test_first','test_last', NOW());`)
	s.assert.Nil(err)

	customer, err := s.store.FindCustomerByEmail(s.store.db, "test_email")
	s.assert.Nil(err)
	s.assert.NotNil(customer)
	s.assert.Equal(customer.Email, "test_email")
}

func (s *StoreSuite) TestFindCustomerByID() {
	_, err := s.store.db.Exec(`INSERT INTO customers (email, first_name, last_name, birth_date) VALUES ('test','test','test', NOW());`)
	s.assert.Nil(err)

	var id int
	err = s.store.db.QueryRow(`SELECT customer_id FROM Customers WHERE email='test';`).
		Scan(&id)

	customer, err := s.store.FindCustomerByID(s.store.db, id)
	s.assert.Nil(err)
	s.assert.NotNil(customer)
	s.assert.Equal(customer.ID, id)
}

func (s *StoreSuite) TestAllCustomers() {
	_, err := s.store.db.Exec(`INSERT INTO customers (email, first_name, last_name, birth_date) VALUES ('test','test','test', NOW());`)
	s.assert.Nil(err)

	customers, err := s.store.AllCustomers(s.store.db)
	s.assert.Nil(err)
	s.assert.Equal(len(customers), 1)
	s.assert.Equal(customers[0].Email, "test")
}
