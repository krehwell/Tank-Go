package database

import (
	"final-project/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

const (
	host = "localhost"
	user = "postgres"
	// password=gorm
	dbname  = "golangFinalProject"
	port    = 5432
	sslmode = "disable"
)

func InitializeDb() Database {
	conStr := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=%s", host, user, dbname, port, sslmode)

	gormDb, gormErr := gorm.Open(postgres.Open(conStr), &gorm.Config{})
	if gormErr != nil {
		fmt.Println("Gorm db error in connecting", gormErr)
		os.Exit(-1)
	}

	if migrateErr := gormDb.AutoMigrate(&model.User{}, &model.Comment{}, &model.Photo{}, &model.SocialMedia{}); migrateErr != nil {
		fmt.Println("Gorm failed to migrate", migrateErr)
		os.Exit(-1)
	}

	fmt.Println("Db initialzed!")
	return Database{gormDb}
}
