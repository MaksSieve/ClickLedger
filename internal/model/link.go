package model

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Target   string `gorm:"not null"`
	LinkType string `gorm:"not null"`
	Slug     string `gorm:"not null"`
}
