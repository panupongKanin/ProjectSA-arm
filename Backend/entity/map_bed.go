package entity

import (

	"gorm.io/gorm"
	"time"
)

type Map_Bed struct {
	gorm.Model
	// Map_Bed_ID		uint		`gorm:"primaryKey;autoIncrement:true"`

	Triage_ID		string
	Triage		Triage	`gorm:"references:id"`

	Admidtime		time.Time

	Bed_ID		string
	Bed			Bed		`gorm:"references:id"`

	Map_Bed_Comment	string

	User_ID		string
	User			User		`gorm:"references:id"`

}