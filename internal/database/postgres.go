package database

import (
	"database/sql"
	"log"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error{

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env variable")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Printf("database name = %s \n",os.Getenv("DB_NAME"))
	
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		return err
	}
	if err = db.Ping(); err != nil{
		log.Fatal("Could not ping the database")
	}

	fmt.Println("you have connected to the database successfully")
	
	DB = db
	return nil
}
