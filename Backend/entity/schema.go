package entity

import(

	"gorm.io/gorm"
	"time"
)

type Zone struct{
	gorm.Model
	Zone_Name string
	Beds []Bed `gorm:"ForeignKey:Zone_ID"`
}

type Bed struct{
	gorm.Model
	Bed_Name string
	Bed_State int
	Zone_ID *int
	Zone Zone `gorm:"references:id"`
	Map_Bed []Map_Bed `gorm:"ForeignKey:Bed_ID"`
}

type Map_Bed struct{
	gorm.Model
	// Triage_ID ทำหน้าที่เป็น FK
	Triage_ID 		*uint
	Triage   	 	Triage 	`gorm:"references:id"`
	Admidtime		time.Time
	// Bed_ID ทำหน้าที่เป็น FK
	Bed_ID 		*uint
	Bed    		Bed		 `gorm:"references:id"`
	MapBed_Comment	string
	// User_ID ทำหน้าที่เป็น FK
	// User_ID 		*uint
	// User    		User 		`gorm:"references:id"`

}

type Gender struct {
	gorm.Model
	Gender_Name string
	Patient []Patient `gorm:"foreignKey:GenderID"`
}
type Patient struct {
	gorm.Model
	Patient_Name   string
	GenderID       *int
	Gender         Gender `gorm:"references:id"`
}

type DiseaseType struct {
	gorm.Model
	DiseaseType_NAME string
	Disease []Disease `gorm:"ForeignKey:DiseaseType_ID"`
}

type Disease struct {
	gorm.Model
	Disease_NAME   string
	DiseaseType_ID *uint
	DiseaseType DiseaseType `gorm:"references:id"` //เพิ่ม `gorm:"references:id"`

	Triages []Triage `grom:"foreignKey:Disease_ID"`
}

type InpantientDepartment struct {
	gorm.Model
	InpantientDepartment_NAME string
	Triages                   []Triage `grom:"foreignKey:IPD_ID"`
}

type Triage struct {
	gorm.Model
	Patient_ID              *uint
	Patient                 Patient
	Disease_ID              *uint
	Disease                 Disease
	InpantientDepartment_ID *uint
	InpantientDepartment    InpantientDepartment
	Triage_COMMENT          string
	Triage_State int						//หญิงเพิ่ม
	Map_Bed []Map_Bed `gorm:"ForeignKey:Triage_ID"` //หญิงเพิ่ม
} 