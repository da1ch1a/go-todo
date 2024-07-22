package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// 疎通確認のため一時的にpublicに
var Db *sql.DB

func InitDb() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
			User:   os.Getenv("DB_USER"),
			Passwd: os.Getenv("DB_PASSWORD"),
			Net:    "tcp",
			Addr:   "127.0.0.1:33060",
			DBName: "go_todo_local",
			AllowNativePasswords: true, // rootでログインするため
	}
	// Get a database handle.
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
			log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
			log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return Db
}