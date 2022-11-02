package generics

import "final-project/database"

type IRepository struct {
	database database.Database
}

type IService struct {
	DB database.Database
}
