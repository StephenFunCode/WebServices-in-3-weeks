// to run postgres locally, you can use docker
// docker pull postgres
// docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5431:5431 -d postgres
//
// otherwise use the sqlite provided in the repo

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully connected!")
}
