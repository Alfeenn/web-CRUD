package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golangdb?parseTime=true")

	fmt.Println("\tSuccessfully connected!")
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(15)
	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute) //reset connection
	return db, err
}
