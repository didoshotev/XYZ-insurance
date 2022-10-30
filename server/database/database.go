package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DbConn *sql.DB

func SetupDatabase() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		fmt.Println("Did not find .env file, refer to .env-example")
		os.Exit(1)
	}

	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "insurancedb",
	}

	var err error
	DbConn, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := DbConn.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to: ", cfg.Addr)
}
