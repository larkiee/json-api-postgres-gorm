package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	UserID string `gorm:"primaryKey;uniqueIndex" json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Addresses []Address `gorm:"foreignKey:UserID;references:UserID" json:"addresses"`
}

type Address struct {
	gorm.Model `json:"-"`
	UserID string `json:"-"`
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}