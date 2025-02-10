package models

import "gorm.io/gorm"

type Volunteer struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	RequestID uint `gorm:"not null"`
	Status    string `gorm:"type:enum('Pending', 'Accepted', 'Rejected');default:'Pending'"`
}
