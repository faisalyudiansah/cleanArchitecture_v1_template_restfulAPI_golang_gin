package databases

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Seed() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close()

	sqlFile := "databases/db.sql"
	content, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatal("Failed to read SQL file:", err)
	}

	_, err = db.ExecContext(context.Background(), string(content))
	if err != nil {
		log.Fatal("Failed to execute SQL:", err)
	}

	fmt.Println("SQL executed successfully")
}
