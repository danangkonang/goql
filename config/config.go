package config

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

func Connection(URI string) *DB {
	arr := strings.Split(URI, "://")
	driverName := arr[0]
	// 	connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
	// 		os.Getenv("DB_USER"),
	// 		os.Getenv("DB_PASSWORD"),
	// 		os.Getenv("DB_HOST"),
	// 		os.Getenv("DB_PORT"),
	// 		os.Getenv("DB_NAME"),
	// 	)
	db, err := sql.Open(driverName, URI)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{DB: db}
}
