package database

import (
	"fmt"
	"log"

	"github.com/bensonopisa/tasktracker/config"
	"github.com/bensonopisa/tasktracker/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s Timezone=Africa/Nairobi",
		dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbName, dbConfig.DbPort, dbConfig.SslMode)

	log.Println("DSN : ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect to the database: ", err)
		return nil, err
	}

	log.Println("Connected to database successfully...Running db migrations")

	err = db.AutoMigrate(&model.Permission{}, &model.Role{}, &model.Task{}, &model.User{})

	if err != nil {
		log.Println("Auto migration failed. Reason:: ", err)
	}

	log.Println("Auto migrations completed successfully.")

	return db, nil
}
