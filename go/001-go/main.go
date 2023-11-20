package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// gin の変数を定義
	router := gin.Default()
	// css ファイルのディレクトリを指定
	router.Static("styles", "./styles")
	// HTML ファイルのディレクトリを指定
	router.LoadHTMLGlob("templates/*.html")

	// "/login" にリダイレクト
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/login")
	})

	// "/login" と "login.html" をマッピング
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})

	// ログインボタンで post
	router.POST("/home", func(c *gin.Context) {
		user := getUserById(c.PostForm("userId"))
		if passwordEncrypt(c.PostForm("password")) == user.Password {
			c.HTML(200, "home.html", gin.H{"name": user.UserName})
		} else {
			c.HTML(200, "login.html", gin.H{})
		}
	})

	// server 起動
	router.Run()
}

// sha256 でパスワードのハッシュ化
func passwordEncrypt(password string) string {
	b := []byte(password)
	sha256 := sha256.Sum256(b)
	hashed := hex.EncodeToString(sha256[:])
	return hashed
}

// ユーザIDからユーザマスタのレコードを取得する
func getUserById(userId string) UserMaster {
	// db 接続
	db, err := sql.Open("mysql", "test:test@tcp(localhost:3306)/go_web?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	defer dbmap.Db.Close()

	// Select
	var user []UserMaster
	_, err = dbmap.Select(&user, "select * from user_master where user_id = '"+userId+"'")

	return user[0]
}

type UserMaster struct {
	UserId     string    `db:"user_id"`
	Password   string    `db:"password"`
	UserName   string    `db:"user_name"`
	CreateDate time.Time `db:"create_date"`
}
