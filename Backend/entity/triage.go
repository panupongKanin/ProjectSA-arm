package entity

import (
	"gorm.io/gorm"
)

type Triage struct {
	gorm.Model
	Triage_ID		uint		`gorm:"primaryKey;autoIncrement:true"`

	Patient_ID		string
	Patient		Patient	`gorm:"references:id"`

	Disease_ID	string
	Disease		Disease	`gorm:"references:id"`

	IPD_ID		string
	Ipd			Ipd		`gorm:"references:id"`

	// IPD_ID		string
	// IPD			IPD		`gorm:"references:id"`
	
	Triage_Comment	string
}