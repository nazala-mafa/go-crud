package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name          string `gorm:"type:varchar(256)" json:"name" faker:"sentence"`
	Email         string `gorm:"type:varchar(256);uniqueIndex" json:"email" faker:"email"`
	Password      string `gorm:"type:varchar(256)" json:"password" faker:"sentence"`
	RememberToken string `gorm:"type:varchar(256);null" json:"remember_token" faker:"sentence"`
	AvatarURL     string `gorm:"type:varchar(256);null" json:"avatar_url" faker:"url"`
}
