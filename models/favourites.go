package models

import "gorm.io/gorm"

type Favourite struct {
	gorm.Model
	UserID uint `gorm:"not null;index" json:"userId"`
	BookID uint `gorm:"not null;index" json:"bookId"`
}
