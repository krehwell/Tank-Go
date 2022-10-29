package users

import (
	"final-project/database"
	"final-project/model"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	database database.Database
}

func (ur *UserRepository) getUserById(id string) (model.User, bool) {
	u := model.User{Id: id}
	if err := ur.database.DB.First(&u, "id = ?", id).Error; err != nil {
		return u, false
	}

	return u, true
}

func (ur *UserRepository) getUserByUsername(email string) (model.User, bool) {
	u := model.User{Email: email}
	if err := ur.database.DB.First(&u, "email = ?", email).Error; err != nil {
		return u, false
	}

	return u, true
}

func (ur *UserRepository) createNewUser(newUser model.User) (model.User, error) {
	createUserErr := ur.database.DB.Create(&newUser).Error
	if createUserErr != nil {
		return model.User{}, createUserErr
	}

	return newUser, nil
}

func (ur *UserRepository) updateUserData(oldUserData, newUserData model.User) (model.User, error) {
	userBuffer := model.User{}
	updateUserErr := ur.database.DB.Model(&userBuffer).
		Clauses(clause.Returning{}).
		Where("Id = ?", oldUserData.Id).
		Updates(&newUserData).Error

	if updateUserErr != nil {
		return model.User{}, updateUserErr
	}

	return userBuffer, nil
}
