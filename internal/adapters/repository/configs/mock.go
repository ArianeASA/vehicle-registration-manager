package configs

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

// MockDatabaseConfig is a mock implementation of the DatabaseConfig
type MockDatabaseConfig struct {
	mock.Mock
}

func (m *MockDatabaseConfig) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDatabaseConfig) InitDatabase() (*sql.DB, error) {
	args := m.Called()
	return args.Get(0).(*sql.DB), args.Error(1)
}

func (m *MockDatabaseConfig) GetDB() *sql.DB {
	args := m.Called()
	return args.Get(0).(*sql.DB)
}

func (m *MockDatabaseConfig) Close() error {
	args := m.Called()
	return args.Error(0)
}
