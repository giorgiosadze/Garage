package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./visits.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS visit_count (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		count INTEGER
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize count
	_, err = db.Exec("INSERT INTO visit_count (count) VALUES (0);")
	if err != nil {
		log.Fatal(err)
	}
}

func countVisitHandler(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("UPDATE visit_count SET count = count + 1 WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}

	var count int
	err = db.QueryRow("SELECT count FROM visit_count WHERE id = 1").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Visit count: %d\n", count)
}

func main() {
	initDB()
	http.HandleFunc("/visit", countVisitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
