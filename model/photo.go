package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model `valid:"-"`
	Id         string `gorm:"primaryKey" valid:"uuid"`
	UserId     string `valid:"uuid"`
	Title      string `gorm:"not null" valid:"-"`
	Caption    string `valid:"-"`
	PhotoUrl   string `gorm:"not null" valid:"url"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewString()
	p.Id = id

	_, validateErr := govalidator.ValidateStruct(p)
	if validateErr != nil {
		return validateErr
	}

	return
}

func (p *Photo) AfterSave(tx *gorm.DB) (err error) {
	_, validateErr := govalidator.ValidateStruct(p)
	if validateErr != nil {
		return validateErr
	}
	return
}
