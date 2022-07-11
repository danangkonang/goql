package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type DB struct {
	DB *sql.DB
}

func Connection() *DB {
	var connection string
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		connection = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
	case "mysql":
		connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	}
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connection)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{DB: db}
}
