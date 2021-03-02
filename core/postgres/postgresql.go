package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	database = &gorm.DB{}
)

// Configuration configuration postgresql
type Configuration struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

// NewConnection initialize a new db connection.
func NewConnection(config Configuration) (err error) {
	postgreSQLCredentials := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		config.Host, config.Port, config.Username, config.Password, config.DatabaseName,
	)

	database, err = gorm.Open(postgres.Open(postgreSQLCredentials), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	return nil

}

// GetDatabase get database
func GetDatabase() *gorm.DB {
	return database
}
