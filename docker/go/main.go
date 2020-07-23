package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	ID        string    `db:"id"`
	Image     string    `db:"image"`
	Name      string    `db:"name"`
	Message   string    `db:"message"`
	Slack     string    `db:"slack"`
	Twitter   string    `db:"twitter"`
	Github    string    `db:"github"`
	GradeID   int       `db:"grade_id"`
	MajorID   int       `db:"major_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT name FROM members")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var member Member
		if err := rows.Scan(&member.Name); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, member.Name)
		fmt.Fprintf(w, "こんにちは")
		fmt.Println(member.Name)
		fmt.Println("こんにちは")
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Printf("Starting server at 'http://localhost:8082/'\n")

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/skillset")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, db)
	})
	http.ListenAndServe(":8082", nil)
}
