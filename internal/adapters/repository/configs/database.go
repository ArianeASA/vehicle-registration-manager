package configs

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

type databaseConfig struct {
	host       string
	port       string
	user       string
	password   string
	dBName     string
	sslMode    string
	driverName string
	schema     string
	db         *sql.DB
}

type DatabaseConfigs interface {
	InitDatabase() (*sql.DB, error)
	GetDB() *sql.DB
	Close() error
	Ping() error
}

func NewDatabaseConfig() DatabaseConfigs {
	return &databaseConfig{
		host:       os.Getenv("DB_HOST"),
		port:       os.Getenv("DB_PORT"),
		user:       os.Getenv("DB_USER"),
		password:   os.Getenv("DB_PASSWORD"),
		dBName:     os.Getenv("DB_NAME"),
		driverName: os.Getenv("DRIVER_NAME"),
		schema:     os.Getenv("DB_SCHEMA"),
		sslMode:    "disable",
	}
}

func (c *databaseConfig) dataSourceName() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		c.host, c.port, c.user, c.password, c.dBName, c.sslMode, c.schema)
}

func (c *databaseConfig) isProd() bool {
	return os.Getenv("SCOPE") == "prod"
}

func (c *databaseConfig) InitDatabase() (*sql.DB, error) {
	if c.isProd() {
		db, err := sql.Open(c.driverName, c.dataSourceName())
		if err != nil {
			return nil, err
		}
		c.db = db
		err = db.Ping()
		return db, err
	}
	return nil, nil
}

func (c *databaseConfig) GetDB() *sql.DB {
	return c.db
}

func (c *databaseConfig) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func (c *databaseConfig) Ping() error {
	if c.isProd() {
		if c.db != nil {
			return c.db.Ping()
		}
		return errors.New("database not initialized")
	}
	return nil
}
