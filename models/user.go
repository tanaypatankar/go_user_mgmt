package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Id         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Country    string `json:"country"`
	Status     string `json:"status"`
}
