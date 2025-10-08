package models

import "time"

type User struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;size:200;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Avatar    string    `gorm:"size:100;null" json:"avatar"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type UserPayload struct {
	Name     string `json:"name" form:"name" binding:"required,min=3"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type UserUpdatePayload struct {
	Name     string `json:"name" form:"name" binding:"omitempty,required,min=3"`
	Email    string `json:"email" form:"email" binding:"omitempty,required,email"`
	Password string `json:"password" form:"password" binding:"omitempty,required,min=6"`
}
