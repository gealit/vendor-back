package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicId uuid.UUID `gorm:"type:uuid" json:"public_id"`
	UserName string    `json:"username"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Role     string    `json:"role"`
	IsActive bool      `gorm:"default:true"`
	Profile  Profile   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// BeforeCreate hook to set UUID if not already set
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.PublicId == uuid.Nil {
		u.PublicId = uuid.New()
	}
	return
}

type Profile struct {
	gorm.Model
	UserID     uint   `gorm:"uniqueIndex;not null" json:"user_id"`
	Name       string `gorm:"size:100" json:"name"`
	LastName   string `gorm:"size:100" json:"lastname"`
	MiddleName string `gorm:"size:100" json:"middlename"`
	Bio        string `gorm:"size:999" json:"bio"`
	Phone      string `gorm:"size:100" json:"phone"`
	Avatar     string `gorm:"size:255" json:"avatar_picture"`
}
