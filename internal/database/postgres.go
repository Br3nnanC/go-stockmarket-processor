package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func connectDB() error{
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")
	)
	fmt.Println("database name = ",os.Getenv("DB_NAME"))
	
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		return err
	}
	if err = db.Ping(); err != nil{
		return err
	}
	fmt.Println("you have connected to the database successfully")
	DB = db
	return nil
}
