package entity

import (
	"gorm.io/gorm"
)

type User_Type struct {
	gorm.Model
	//User_Type_ID 		uint 		`gorm:"primaryKey;autoIncrement:true"`
	UserType	 			string
}
