package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DBConfig holds database configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// GetDefaultDBConfig returns the default database configuration
func GetDefaultDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "minho",
		Password: "1234",
		DBName:   "smartOrder",
	}
}

// NewDBFromConfig creates a new database connection from configuration
func NewDBFromConfig(config *DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	// 데이터베이스 연결
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	// 데이터베이스 연결 테스트
	err = db.Ping()
	if err != nil {
		db.Close() // 연결 실패시 리소스 정리
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Successfully connected to database")
	return db, nil
}

// InitDB initializes database with default configuration
func InitDB() (*sql.DB, error) {
	return NewDBFromConfig(GetDefaultDBConfig())
}
