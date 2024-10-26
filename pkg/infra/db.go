package infra

import (
	"da1ch1a/go-todo/pkg/core"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

// 疎通確認のため一時的にpublicに
var Db *sql.DB

func InitDb() *sql.DB {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:33060",
		DBName:               "go_todo_local",
		AllowNativePasswords: true, // rootでログインするため
		Params: map[string]string{
			// MySQLのDATEとDATETIMEはGoで出力するとデフォルトは[]byteになるので、これをtime.Timeに変換するためにparseTime=trueを指定する
			// NOTE: https://qiita.com/stern327/items/cb325b44a0335deea402
			"parseTime": "true",
			"loc":       "UTC",
		},
	}
	// Get a database handle.
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	Db = sqldblogger.OpenDriver(
		cfg.FormatDSN(),
		Db.Driver(),
		zerologadapter.New(core.Logger),
		sqldblogger.WithMinimumLevel(sqldblogger.LevelDebug),
	)

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return Db
}
