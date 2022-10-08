package entity

import (

	"gorm.io/gorm"
)

type Triage struct {
	gorm.Model
	Triage_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Patient_ID		uint
	Disease_ID		uint
	IPD_ID		uint
	Triage_Comment	string
	UserID		uint
}