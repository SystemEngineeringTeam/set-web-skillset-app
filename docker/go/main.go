package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	// メンバーのデータと、member_idで紐づいた各種idをそれぞれ「,」で連結した文字列にして取得
	memberRows, err := db.Query("SELECT a.*, GROUP_CONCAT( b.group_id SEPARATOR ',' ), GROUP_CONCAT( c.tech_area_id SEPARATOR ',' ) FROM members a LEFT JOIN groups_of_members b on a.id = b.member_id LEFT JOIN tech_areas_of_members c on a.id = c.member_id GROUP BY a.id;")
	if err != nil {
		log.Fatal(err)
	}
	defer memberRows.Close()

	var response []ResponseMember

	for memberRows.Next() {
		var member Member
		var groups string
		var techAreas string
		if err := memberRows.Scan(&member.ID, &member.Image, &member.Name, &member.Message, &member.Slack, &member.Twitter, &member.Github, &member.GradeID, &member.MajorID, &member.CreatedAt, &member.UpdatedAt, &groups, &techAreas); err != nil {
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
		//連結して取得した文字列を分割して配列に格納
		groupsArray := strings.Split(groups, ",")
		for _, group := range groupsArray {
			g, err := strconv.Atoi(group)
			if err != nil {
				log.Fatal(err)
			}
			responseMember.Group = append(responseMember.Group, g)
		}
		techAreasArray := strings.Split(techAreas, ",")
		for _, techArea := range techAreasArray {
			t, err := strconv.Atoi(techArea)
			if err != nil {
				log.Fatal(err)
			}
			responseMember.TechnicalArea = append(responseMember.TechnicalArea, t)
		}

		//response用の配列に構造体を追加
		response = append(response, responseMember)
	}
	if err := memberRows.Err(); err != nil {
		log.Fatal(err)
	}

	//response用の配列をjsonに整形
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
