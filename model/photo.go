package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model `valid:"-"`
	Id         string     `gorm:"primaryKey" valid:"uuid"`
	UserId     string     `valid:"uuid"`
	Title      string     `gorm:"not null" valid:"-"`
	Caption    string     `valid:"-"`
	PhotoUrl   string     `gorm:"not null" valid:"url"`
	IsDeleted  bool       `gorm:"type:boolean;default:false" valid:"-"`
	Comments   []*Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" valid:"-"`
}

type PhotoBody struct {
	Id       string
	Title    string
	Caption  string
	PhotoUrl string
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
