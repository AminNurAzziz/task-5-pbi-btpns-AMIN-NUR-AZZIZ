package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Title    string `gorm:"not null"`
	Caption  string
	PhotoUrl string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
