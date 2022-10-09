package entity

import (
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	//Gender_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Gender_Type		string
}