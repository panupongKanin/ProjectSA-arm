package entity

import (

	"gorm.io/gorm"
)

type Bed struct {
	gorm.Model
	//Bed_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Bed_Name		string

	Zone_ID		string
	Zone			Zone		`gorm:"references:id"`
	
}