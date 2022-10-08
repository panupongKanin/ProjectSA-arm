package entity

import (

	"gorm.io/gorm"
)

type userType struct {
	gorm.Model
	Usertype_ID 		uint 		`gorm:"primaryKey;autoIncrement:true"`
	Usertype 			string
}
