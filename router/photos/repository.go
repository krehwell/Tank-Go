package photos

import (
	"final-project/database"
	"final-project/model"
)

type PhotoRepository struct {
	database database.Database
}

func (pr *PhotoRepository) createPhoto(newPhotoData model.Photo) (model.Photo, error) {
	createUserErr := pr.database.DB.Create(&newPhotoData).Error
	if createUserErr != nil {
		return model.Photo{}, createUserErr
	}

	return newPhotoData, nil
}

func (pr *PhotoRepository) getPhotosByUserId(userId string) ([]model.Photo, error) {
	photosBuffer := []model.Photo{}

	if err := pr.database.DB.Where("user_id = ?", userId).Find(&photosBuffer).Error; err != nil {
		return []model.Photo{}, err
	}

	return photosBuffer, nil
}
