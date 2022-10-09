package entity

import (
	"gorm.io/gorm"
)

type Ipd struct {
	gorm.Model
	//IPD_ID		uint			`gorm:"primaryKey;autoIncrement:true"`
	IPD_Name		string

}