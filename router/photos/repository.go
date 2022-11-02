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
