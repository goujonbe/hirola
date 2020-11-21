package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var connectionString string
	var csvPath string

	flag.StringVar(&csvPath, "csv-path", "", "absolute csv path")
	flag.StringVar(&connectionString, "to", "", "db url")
	flag.Parse()

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("COPY test FROM '/Users/benoit/go/src/github.com/goujonbe/hirola/transactions.csv' csv header;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}
