package entity

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	//Patient_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Patient_Name	string

	Gender_ID		string
	Gender		Gender	 `gorm:"references:id"`	

}