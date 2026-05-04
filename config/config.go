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
	var u string
	switch driverName {
	case "postgres":
		u = URI
	case "mysql":
		u = arr[1]
	}
	db, err := sql.Open(driverName, u)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{DB: db}
}
