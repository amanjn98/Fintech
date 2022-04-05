package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)
// The "db" package level variable will hold the reference to our database instance
var db *sql.DB
func initDB(){
	var err error
	// Connect to the postgres db
	dbhost,dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbhost, 5432, dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	log.Println("Database connection established")
}

func GetDB() *sql.DB {
	if db ==nil{
		initDB()
	}
	return db
}


