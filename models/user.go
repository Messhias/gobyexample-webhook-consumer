package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ExternalID uuid.UUID `gorm:"primaryKey;unique;type:uuid;default:(gen_random_uuid())"`
	Email      string    `gorm:"unique"`
	Active     bool      `gorm:"default:true"`
}
