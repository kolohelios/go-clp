package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// packages being imported only for their side-effects:
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		dbUsername = os.Getenv("GO_CLP_DB_USERNAME")
		dbPassword = os.Getenv("GO_CLP_DB_PASSWORD")
		dbURL      = os.Getenv("GO_CLP_DB_URL")
		dbName     = os.Getenv("GO_CLP_DB_NAME")
	)

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUsername, dbPassword, dbURL, dbName)

	db, err := sql.Open("mariadb", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
