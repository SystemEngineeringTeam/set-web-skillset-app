package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Member 名簿に登録されている会員情報一人分を表す型
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

type product struct {
	Overview string `json:"overview"`
	URL      string `json:"url"`
}

type ResponseMember struct {
	ID            string    `json:"id"`
	Image         string    `json:"image"`
	Name          string    `json:"name"`
	Message       string    `json:"message"`
	Slack         string    `json:"slack"`
	Twitter       string    `json:"twitter"`
	Github        string    `json:"github"`
	GradeID       int       `json:"grade_id"`
	MajorID       int       `json:"major_id"`
	Group         []int     `json:"group"`
	TechnicalArea []int     `json:"technical_area"`
	Technology    []int     `json:"technology"`
	Product       []product `json:"product"`
}

func get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM members")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var response []ResponseMember

	for rows.Next() {
		var member Member
		if err := rows.Scan(&member.ID, &member.Image, &member.Name, &member.Message, &member.Slack, &member.Twitter, &member.Github, &member.GradeID, &member.MajorID, &member.CreatedAt, &member.UpdatedAt); err != nil {
			log.Fatal(err)
		}

		var responseMember ResponseMember
		responseMember.ID = member.ID
		responseMember.Image = member.Image
		responseMember.Name = member.Name
		responseMember.Message = member.Message
		responseMember.Slack = member.Slack
		responseMember.Twitter = member.Twitter
		responseMember.Github = member.Github
		responseMember.GradeID = member.GradeID
		responseMember.MajorID = member.MajorID

		response = append(response, responseMember)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	responseJSON, err := json.Marshal(&response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(responseJSON))
}

func main() {
	fmt.Printf("Starting server at 'http://localhost:8082/'\n")
	fmt.Printf("goto'http://localhost:8082/member/get'\n")

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/skillset?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/member/get", func(w http.ResponseWriter, r *http.Request) {
		get(w, r, db)
	})
	http.ListenAndServe(":8082", nil)
}
