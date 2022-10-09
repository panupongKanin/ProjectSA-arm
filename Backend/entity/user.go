package entity

import (
	"gorm.io/gorm"
)



type User struct {
	gorm.Model
	//User_ID		uint			`gorm:"primaryKey;autoIncrement:true"`
	User_Name		string
	User_Password	string

	User_Type_ID	string
	User_Type		User_Type		 `gorm:"references:id"`	

	Name			string
	
}

