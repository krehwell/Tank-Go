package model

import (
	"final-project/utils"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model  `valid:"-"`
	Id          string      `gorm:"primaryKey" valid:"uuid"`
	Username    string      `gorm:"unique;not null" valid:"stringlength(3|20)"`
	Email       string      `gorm:"unique;not null" valid:"email"`
	Password    string      `gorm:"not null" valid:"stringlength(6|200)"`
	Age         int         `gorm:"not null" valid:"range(8|200)"`
	IsDeleted   bool        `gorm:"type:boolean;default:false" valid:"-"`
	Photo       Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"-"`
	Comment     Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"-"`
	SocialMedia SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewString()
	u.Id = id

	_, validateErr := govalidator.ValidateStruct(u)
	if validateErr != nil {
		return validateErr
	}

	hashedPassword := utils.HashAndSalt([]byte(u.Password))
	u.Password = hashedPassword

	return
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	_, validateErr := govalidator.ValidateStruct(u)
	if validateErr != nil {
		return validateErr
	}
	return
}
