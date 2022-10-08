package entity

import (

	"gorm.io/gorm"
)

type Bed struct {

	gorm.Model
	Bed_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Bed_Name		string
	Zone_ID		Zone		`gorm:"references:Zone_ID"`
	
}