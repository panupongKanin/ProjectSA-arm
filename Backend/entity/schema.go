package entity

import (
	// "time"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	User_Name     string `gorm:"uniqueIndex"`
	User_Password string
	Name          string

	// User_Type_ID  ทำหน้าที่เป็น FK
	User_Type_ID *uint
	// เป็นข้อมูล user_type เมื่อ join ตาราง
	User_Type User_Type `gorm:"references:id"`
}

type User_Type struct {
	gorm.Model

	UserType string

	// 1 User_Type เป็นเจ้าของได้หลาย Users
	Users []User `gorm:"foreignKey:User_Type_ID"`
}

type Zone struct {
	gorm.Model

	Zone_Name string `gorm:"uniqueIndex"`

	// 1 Zone เป็นเจ้าของได้หลาย Beds
	Beds []Bed `gorm:"foreignKey:Zone_ID"`
}

type Bed struct {
	gorm.Model

	Bed_Name string 

	// Zone_ID   ทำหน้าที่เป็น FK
	Zone_ID *uint
	// เป็นข้อมูล Zone เมื่อ join ตาราง
	Zone Zone `gorm:"references:id"`
}

// =================================== ตารางหลัก ===================================
// =================================== ตารางหลัก ===================================
type Map_Bed struct {
	gorm.Model

	// Triage_ID ทำหน้าที่เป็น FK
	Triage_ID 		*uint
	Triage   		Triage 	`gorm:"references:id" valid:"-"` 

	Admidtime		time.Time

	// Bed_ID ทำหน้าที่เป็น FK
	Bed_ID 		*uint
	Bed	  		Bed 	`gorm:"references:id" valid:"-"` 

	Triage_Comment	string

	// User_ID ทำหน้าที่เป็น FK
	User_ID 		*uint
	User  		User 		`gorm:"references:id" valid:"-"`
}
// =================================== ตารางหลัก ===================================
// =================================== ตารางหลัก ===================================

// =================================== ดึงมาจากเพื่อน ===================================

type Disease struct {
	gorm.Model
	//Disease_ID		uint			`gorm:"primaryKey;autoIncrement:true"`
	Disease_Name string

	// Disease_Type_ID   ทำหน้าที่เป็น FK
	Disease_Type_ID *uint
	// เป็นข้อมูล Zone เมื่อ join ตาราง
	Disease_Type Disease_Type `gorm:"references:id"`
}

type Disease_Type struct {
	gorm.Model

	DiseaseType string

	Diseases []Disease `gorm:"foreignKey:Disease_Type_ID"`
}

type Ipd struct {
	gorm.Model

	IPD_Name string
}

type Patient struct {
	gorm.Model
	//Patient_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Patient_Name string

	// Gender_ID   ทำหน้าที่เป็น FK
	Gender_ID *uint
	// เป็นข้อมูล Gender เมื่อ join ตาราง
	Gender Gender `gorm:"references:id"`
}

type Gender struct {
	gorm.Model
	//Gender_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Gender_Type string

	Patients []Patient `gorm:"foreignKey:Gender_ID"`
}

type Triage struct {
	gorm.Model

	// Patient_ID ทำหน้าที่เป็น FK
	Patient_ID *uint
	Patient    Patient `gorm:"references:id" valid:"-"`

	// Disease_ID ทำหน้าที่เป็น FK
	Disease_ID *uint
	Disease    Disease `gorm:"references:id" valid:"-"`

	// IPD_ID ทำหน้าที่เป็น FK
	IPD_ID         *uint
	Ipd            Ipd `gorm:"references:id" valid:"-"`
	Triage_Comment string
}

// =================================== ดึงมาจากเพื่อน ===================================
