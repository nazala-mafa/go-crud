package models

import "gorm.io/gorm"

type RevokedToken struct {
	gorm.Model
	Token string `gorm:"type:varchar(512);uniqueIndex"`
}
