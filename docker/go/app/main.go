package main

import (
	"log"
	"net/http"
	"database/sql"
	// "fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/labstack/echo"
)

// type User struct {
//     User_ID int `db:"user_id"`
//     Name string `db:"user_name"`
// 	email string `db:"email"`
//     Password  string `db:"Password"`
// }
// type Userlist []User

func main() {

	//データベース：issにアクセス
	db, err := sql.Open("mysql", "root:root@tcp(iss-mysql:3306)/iss")

	err = db.Ping()
    if err != nil {
		//アクセスできない場合はここでエラーがでる
        panic(err)
    }
	defer db.Close()

	// _, err = db.Exec(
    // `INSERT INTO "User" ("user_id", "user_email", "user_name" , "password") VALUES (?, ?, ?, ?) `,
    // 1,
	// "a@gmail.com",
	// "ken",
	// "pass",
	// )

	// if err != nil {
    // 	return err
	// }   

	r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:8080"},
        AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"*"},
    }))

    r.GET("/test", test.get_test)
    r.Run(":8082")
	

	//とりあえず起動してlocalhost:8082にアクセスできるか確認用
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":8082"))
	// e.Startの中はdocker-composeのgoコンテナで設定したportsを指定してください。
}
