package database

import (
	"fmt"
	"log"
	"os"
	"selfregi/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func EntOpen() (*ent.Client, error) {
	err := godotenv.Load((".env"))
    if err != nil {
        panic ("failed reading env file")
    }
    var USER string
	var PASS string
	var PROTOCOL string
	var DBNAME string

	env := os.Getenv("ENV")
    if env == "dev" {
		USER = os.Getenv("DB_USER")
		PASS = os.Getenv("DB_PASS")
		DBNAME = os.Getenv("DB_DATABASE")
		PROTOCOL = os.Getenv("DB_PROTOCOL")
	}
    client := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := ent.Open("mysql", client)
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    } else {
        fmt.Println("database connect ok")
    }
    return db, err
}