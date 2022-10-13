package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `valid:"-"`
	Id       string `gorm:"primaryKey" valid:"uuid"`
	Username string `gorm:"unique;not null" valid:"stringlength(3|20)"`
	Email    string `gorm:"unique;not null" valid:"email"`
	Password string `gorm:"not null" valid:"stringlength(6|200)"`
	Age      int    `gorm:"unique;not null" valid:"range(8|200)"`
}

type Photo struct {
	gorm.Model `valid:"-"`
	Id       string `gorm:"primaryKey" valid:"uuid"`
	Title    string `gorm:"not null" valid:"-"`
	Caption  string `valid:"-"`
	PhotoUrl string `gorm:"not null" valid:"url"`
}

type Comment struct {
	gorm.Model `valid:"-"`
	UserId  string `valid:"uuid"`
	PhotoId string `valid:"-"`
	Message string `gorm:"not null" valid:"-"`
}

type SocialMedia struct {
	Id             string `gorm:"primaryKey" valid:"uuid"`
	Name           string `gorm:"type:varchar(255);not null" valid:"-"`
	SocialMediaUrl string `gorm:"not null" valid:"-"`
	UserId         string `valid:"-"`
}
