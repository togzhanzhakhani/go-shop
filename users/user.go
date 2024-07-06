package users

import (
	"time"
)

type User struct {
	ID               uint           `gorm:"primaryKey"`
	Name             string         `gorm:"not null" validate:"required"`
	Email            string         `gorm:"unique;not null" validate:"required,email"`
	Address          string         `gorm:"not null" validate:"required"`
	RegistrationDate time.Time      `gorm:"not null;autoCreateTime"`
	Role             string         `gorm:"not null" validate:"required,oneof=administrator client"`
}

var UserBaseMessages = map[string]string{
	"required": "is required",
	"email":    "is not valid",
	"oneof":    "must be either 'admin' or 'client'",
}

func (User) TableName() string {
    return "project2_users"
}
