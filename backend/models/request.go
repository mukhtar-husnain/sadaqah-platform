package models

import "gorm.io/gorm"

type VolunteerRequest struct {
	gorm.Model
	UserID             uint   `gorm:"not null"`
	Title              string `gorm:"not null"`
	Description        string `gorm:"not null"`
	Location           string
	EventDate          string `gorm:"not null"`
	RequiredVolunteers int    `gorm:"not null"`
	SkillsRequired     string
}
