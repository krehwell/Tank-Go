package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model `valid:"-"`
	Id         string `gorm:"primaryKey" valid:"uuid"`
	UserId     string `valid:"uuid"`
	PhotoId    string `valid:"-"`
	Message    string `gorm:"not null" valid:"-"`
}

type SocialMedia struct {
	Id             string `gorm:"primaryKey" valid:"uuid"`
	Name           string `gorm:"type:varchar(255);not null" valid:"-"`
	SocialMediaUrl string `gorm:"not null" valid:"-"`
	UserId         string `valid:"-"`
}
