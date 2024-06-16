package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DatabaseConnection() *sql.DB {
	flag := "local"
	if flag == "local" {
		return databaseConnectionLocal()
	}
	return databaseConnectionLocal()
}

func databaseConnectionLocal() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Some error occured:", err)
	}

	dbCredentials := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_USER_PW"),
		Net:                  os.Getenv("DB_USER_NET"),
		Addr:                 os.Getenv("DB_USER_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	dbLocal, err := sql.Open("mysql", dbCredentials.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pinErr := dbLocal.Ping()
	if pinErr != nil {
		log.Fatal(pinErr)
	}

	return dbLocal
}
