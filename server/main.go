package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/table-tap/api/db"
)

var (
	DBConn *db.DB
)

func main() {
	err := godotenv.Overload()
	if err != nil {
		log.Println("No .env file found")
	}

	connStr := os.Getenv("TABLE_TAP_DB")
	if connStr == "" {
		panic("NO DB CONNECTION STRING IN env")
	}

	DBConn = openDB()

	port := os.Getenv("PORT")
	log.Printf("port:%s\n", port)
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := GetRootRouter()

	log.Printf("Starting server on port %s\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)

	if err != nil {
		panic(err)
	}
}

func openDB() *db.DB {
	// setup DB
	connStr := os.Getenv("TABLE_TAP_DB")
	if connStr == "" {
		panic("NO DB CONNECTION STRING IN env")
	}

	conn := sqlx.MustConnect("pgx", connStr)

	encryptKey64 := os.Getenv("ENCRYPTION_KEY")
	encryptKey, err := base64.StdEncoding.DecodeString(encryptKey64)
	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(10 * time.Minute)

	return &db.DB{DB: conn, EncryptKey: (*[32]byte)(encryptKey)}
}
