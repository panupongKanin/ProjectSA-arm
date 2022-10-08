package entity

import (

	"gorm.io/gorm"
	"time"
)

type Nap_Bed struct {
	gorm.Model
	Map_Bed_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Triage_ID		uint		`gorm:"references:Triage_ID"`
	Admidtime		time.Time
	Bed_ID		uint		`gorm:"references:Bed_ID"`
	Map_Bed_Comment	string
	UserID		uint		`gorm:"references:UserID"`

}