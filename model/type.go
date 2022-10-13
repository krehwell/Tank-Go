package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null" valid:"email"`
	Password string `gorm:"unique;not null" valid:"stringlength(6|200)"`
	Age      int    `gorm:"unique;not null" valid:"range(8|200)"`
}

type Photo struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	Title    string `gorm:"not null"`
	Caption  string
	PhotoUrl string `gorm:"not null"`
}

type Comment struct {
	gorm.Model
	UserId  string
	PhotoId string
	Message string `gorm:"not null"`
}

type SocialMedia struct {
	Id             string `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(255);not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserId         string
}
