package entity

import (

	"gorm.io/gorm"
)

type Zone struct {
	gorm.Model
	//Zone_ID		uint 		`gorm:"primaryKey;autoIncrement:true"`
	Zone_Name		string	
}