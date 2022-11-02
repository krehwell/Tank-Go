package database

import (
	"final-project/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host = "localhost"
	user = "postgres"
	// password=gorm
	dbname  = "golangFinalProject"
	port    = 5432
	sslmode = "disable"
)

type Database struct {
	DB *gorm.DB
}

var GORM_CUSTOM_LOGGER = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
	},
)

func InitializeDb() Database {
	conStr := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=%s", host, user, dbname, port, sslmode)

	gormDb, gormErr := gorm.Open(postgres.Open(conStr), &gorm.Config{Logger: GORM_CUSTOM_LOGGER})
	if gormErr != nil {
		fmt.Println("Gorm db error in connecting", gormErr)
		os.Exit(-1)
	}

	if migrateErr := gormDb.AutoMigrate(&model.User{}, &model.Comment{}, &model.Photo{}, &model.SocialMedia{}); migrateErr != nil {
		fmt.Println("Gorm failed to migrate", migrateErr)
		os.Exit(-1)
	}

	fmt.Println("Db initialized!")
	return Database{DB: gormDb}
}
