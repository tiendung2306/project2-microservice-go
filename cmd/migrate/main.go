package main

import (
	"flag"
	"log"
	"os"

	"project2-microservice-go/database"

	_ "github.com/joho/godotenv/autoload" // Tự động load .env
)

func main() {
	// Parse command line flags
	migrateFlag := flag.Bool("migrate", false, "Tạo hoặc cập nhật tables trong database")
	flag.Parse()

	// Check if any flags were provided
	if !*migrateFlag {
		log.Println("Không có lệnh nào được chỉ định. Sử dụng --migrate để tạo tables")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Initialize database
	db := database.New()
	defer db.Close()

	// Run migration if specified
	if *migrateFlag {
		log.Println("Start creating tables...")
		if err := db.AutoMigrate(); err != nil {
			log.Fatalf("Creating tables errors: %v", err)
		}
		log.Println("Creating tables successfully!")
	}
}
