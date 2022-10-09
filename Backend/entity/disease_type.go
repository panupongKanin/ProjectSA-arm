package entity

import (
	"gorm.io/gorm"
)

type Disease_Type struct {
	gorm.Model
	//Disease_Type_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	DiseaseType			string
}