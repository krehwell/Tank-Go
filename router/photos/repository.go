package photos

import (
	"final-project/database"
	"final-project/model"

	"gorm.io/gorm/clause"
)

type PhotoRepository struct {
	database database.Database
}

func (pr *PhotoRepository) getPhotoById(id string) (model.Photo, error) {
	photo := model.Photo{}

	if err := pr.database.DB.First(&photo, "id = ? AND is_deleted = ?", id, false).Error; err != nil {
		return model.Photo{}, err
	}

	return photo, nil
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

	if err := pr.database.DB.Where("user_id = ? AND is_deleted = ?", userId, false).Find(&photosBuffer).Error; err != nil {
		return []model.Photo{}, err
	}

	return photosBuffer, nil
}

func (pr *PhotoRepository) updatePhotoData(oldPhotoData, newPhotoData model.Photo) (model.Photo, error) {
	updateUserErr := pr.database.DB.Model(&oldPhotoData).
		Clauses(clause.Returning{}).
		Where("Id = ?", oldPhotoData.Id).
		Updates(&newPhotoData).Error

	if updateUserErr != nil {
		return model.Photo{}, updateUserErr
	}

	return oldPhotoData, nil
}
