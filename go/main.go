package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbConnection *sql.DB

func main() {
	log.Println("Started")
	dbConnection = openDBConnection()

	var query string
	query = "SELECT 123 WHERE ? = 1"
	log.Println(query)
	rows1, err := dbConnection.Query(query, 1)
	if err != nil {
		log.Panicln("Error querying database", err)
	}
	defer rows1.Close()
	printRows(rows1)

	query = "SELECT 123 WHERE 1 = 1"
	log.Println(query)
	rows2, err := dbConnection.Query(query)
	if err != nil {
		log.Panicln("Error querying database", err)
	}
	defer rows2.Close()
	printRows(rows2)

	log.Println("Ready")
}

func printRows(rows *sql.Rows) {
	var field interface{}
	for rows.Next() {
		if err := rows.Scan(&field); err != nil {
			log.Fatalln(err)
		}
		log.Printf("%[1]T: %#[1]v", field)
	}
}

func openDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "db:secret@tcp(db:3306)/db")
	if err != nil {
		log.Fatalln("Cannot connect to database", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("Database connection failed", err)
	}
	return db
}
