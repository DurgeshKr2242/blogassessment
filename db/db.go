package db

import (
	"database/sql"
	"fmt"

	"github.com/DurgeshKr2242/blogassessment/config"
	_ "github.com/lib/pq"
)

// ConnectDB opens a connection to PostgreSQL using the standard library
func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %w", err)
	}
	
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping db: %w", err)
	}

	return db, nil
}
