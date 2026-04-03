package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ExternalID uuid.UUID `gorm:"primaryKey;unique;type:uuid;default:"`
	Email      string    `gorm:"unique"`
	Active     bool      `gorm:"default:true"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.ExternalID = uuid.New()

	return
}
