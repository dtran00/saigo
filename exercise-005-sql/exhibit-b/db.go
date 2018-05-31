package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicOn(err)

	_, err = db.Exec("DELETE FROM people WHERE ssn = $1", 444556666)
	PanicOn(err)

	_, err = db.Exec("UPDATE people SET ssn = $2 WHERE ssn = $1", 111223333, 123456789)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Alice", 555555555)
	PanicOn(err)

	var person_id int
	row := db.QueryRow("SELECT person_id FROM people WHERE ssn = $1", 555555555)
	err = row.Scan(&person_id)
	PanicOn(err)
	fmt.Printf("person_id: %d\n", person_id)
}
