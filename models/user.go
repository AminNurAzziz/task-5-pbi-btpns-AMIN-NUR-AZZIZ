package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null;minLength:6"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// BeforeDelete akan dijalankan sebelum pengguna dihapus
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
    // Hapus semua foto yang terkait dengan pengguna ini
    if err := tx.Model(&Photo{}).Where("user_id = ?", u.ID).Delete(&Photo{}).Error; err != nil {
        return err
    }
    return nil
}