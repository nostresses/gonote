package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	database *gorm.DB
}

// Opens a connection to the database via a DSN link
func New(dataSourceName string) (*Storage, error) {
	connection := mysql.Open(dataSourceName)

	database, err := gorm.Open(connection, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	database.AutoMigrate(&Note{})

	return &Storage{database: database}, nil
}
