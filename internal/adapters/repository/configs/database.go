package configs

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	DBName     string
	SSLMode    string
	DriverName string
	Schema     string
	db         *sql.DB
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DriverName: os.Getenv("DRIVER_NAME"),
		Schema:     os.Getenv("DB_SCHEMA"),
		SSLMode:    "disable",
	}
}

func (c *DatabaseConfig) dataSourceName() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode, c.Schema)
}

func (c *DatabaseConfig) isProd() bool {
	return os.Getenv("SCOPE") == "prod"
}

func (c *DatabaseConfig) InitDatabase() (*sql.DB, error) {
	if c.isProd() {
		db, err := sql.Open(c.DriverName, c.dataSourceName())
		c.db = db
		return db, err
	}
	return nil, nil
}

func (c *DatabaseConfig) GetDB() *sql.DB {
	return c.db
}

func (c *DatabaseConfig) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func (c *DatabaseConfig) Ping() error {
	if c.isProd() {
		if c.db != nil {
			return c.db.Ping()
		}
		return errors.New("database not initialized")
	}
	return nil
}
