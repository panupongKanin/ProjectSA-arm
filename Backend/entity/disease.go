package entity

import (
	"gorm.io/gorm"
)

type Disease struct {
	gorm.Model
	//Disease_ID		uint			`gorm:"primaryKey;autoIncrement:true"`
	Disease_Name	string

	Disease_Type_ID	string
	Disease_type	Disease_Type	`gorm:"references:id"`

}