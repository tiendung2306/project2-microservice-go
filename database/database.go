package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// DB returns the underlying *gorm.DB instance
	DB() *gorm.DB

	AutoMigrate() error
}

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	schema     = os.Getenv("BLUEPRINT_DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	database = os.Getenv("BLUEPRINT_DB_DATABASE")
	password = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username = os.Getenv("BLUEPRINT_DB_USERNAME")
	port = os.Getenv("BLUEPRINT_DB_PORT")
	host = os.Getenv("BLUEPRINT_DB_HOST")
	schema = os.Getenv("BLUEPRINT_DB_SCHEMA")

	// Validate and set default values for environment variables
	if host == "" || username == "" || password == "" || database == "" || port == "" || schema == "" {
		log.Fatalf("One or more required environment variables are not set: host=%s, user=%s, dbname=%s, port=%s, schema=%s", host, username, database, port, schema)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		host, username, password, database, port, schema)

	// Configure GORM logger
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Log SQL slower than 1 second
			LogLevel:                  logger.Info, // Log everything
			IgnoreRecordNotFoundError: true,        // Ignore record not found errors
			ParameterizedQueries:      false,       // Don't log parameters
			Colorful:                  true,        // Colorful output
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// DB returns the underlying GORM DB connection
func (s *service) DB() *gorm.DB {
	return s.db
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	sqlDB, err := s.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("Failed to get database: %v", err)
		log.Printf("Failed to get database: %v", err)
		return stats
	}

	// Ping the database
	err = sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Printf("db down: %v", err)
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 {
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
func (s *service) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	log.Printf("Disconnected from database: %s", database)
	return sqlDB.Close()
}
