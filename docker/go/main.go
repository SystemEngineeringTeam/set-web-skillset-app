package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func test(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var name string
	err := db.QueryRow("SELECT user_name FROM users LIMIT 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	println(name)

	fmt.Fprintf(w, name)
}

func main() {
	fmt.Printf("Starting server at 'localhost:8082'\n")

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/skillset")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		test(w, r, db)
	})
	http.ListenAndServe(":8082", nil)
}
