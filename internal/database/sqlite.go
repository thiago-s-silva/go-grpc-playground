package database

import (
	"os"

	"github.com/thiago-s-silva/grpc-example/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"

	// Check if db already exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// DB Connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// DB Migration
	err = db.AutoMigrate(&entities.Category{}, &entities.Course{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
