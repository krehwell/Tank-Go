package users

import (
	"final-project/database"
	"final-project/model"
)

type UserRepository struct {
	database database.Database
}

func (ur *UserRepository) createNewUser(newUser model.User) (model.User, error) {
	createUserErr := ur.database.DB.Create(&newUser).Error
	if createUserErr != nil {
		return model.User{}, createUserErr
	}

	return newUser, nil
}
