package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// TODO 中身はコピペなので書き換える
func InitDb() {
	// Capture connection properties.
	cfg := mysql.Config{
			User:   os.Getenv("DBUSER"),
			Passwd: os.Getenv("DBPASS"),
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
			log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
			log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
