package models

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	RequestID     uint   `gorm:"not null"`
	GivenByUserID uint   `gorm:"not null"`
	GivenToUserID uint   `gorm:"not null"`
	Rating        int    `gorm:"not null"`
	Comments      string
}
