package entity

import (
	// "time"
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	User_Name     	string `gorm:"uniqueIndex"`
	User_Password 	string
	Name          	string

	// User_Type_ID  ทำหน้าที่เป็น FK
	User_Type_ID 	*uint
	User_Type    	User_Type `gorm:"references:id"`
	
	Map_Beds 		[]Map_Bed `gorm:"ForeignKey:User_ID"`
	
}

type User_Type struct {
	gorm.Model
	UserType string
	Users []User `gorm:"ForeignKey:User_Type_ID"`
}

type Zone struct {
	gorm.Model

	Zone_Name string `gorm:"uniqueIndex"`

	// 1 Zone เป็นเจ้าของได้หลาย Beds
	
	Beds	[]Bed	`gorm:"ForeignKey:Zone_ID"`
}

type Bed struct {
	gorm.Model

	Bed_Name string 
	// Zone_ID   ทำหน้าที่เป็น FK
	Zone_ID      	*int
	Zone    		Zone 		`gorm:"references:id"`

	Map_Beds 		[]Map_Bed 	`gorm:"ForeignKey:Bed_ID"`

}

type Test struct {
	gorm.Model

	Bed_Name string 
	Zone_Name string 

}

// =================================== ตารางหลัก ===================================
// =================================== ตารางหลัก ===================================
type Map_Bed struct {
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
	User_ID 		*uint
	User    		User 		`gorm:"references:id"`

}
// =================================== ตารางหลัก ===================================
// =================================== ตารางหลัก ===================================

// =================================== ดึงมาจากเพื่อน ===================================

type Disease struct {
	gorm.Model
	//Disease_ID		uint			`gorm:"primaryKey;autoIncrement:true"`
	Disease_Name string

	// Disease_Type_ID   ทำหน้าที่เป็น FK
	Disease_Type_ID 	*uint
	Disease_Type    	Disease_Type	`gorm:"references:id"`

	Triages 		[]Triage 		`gorm:"ForeignKey:Disease_ID"`

}

type Disease_Type struct {
	gorm.Model

	DiseaseType string

	Diseases []Disease `gorm:"ForeignKey:Disease_Type_ID"`
}

type Ipd struct {
	gorm.Model
	IPD_Name string

	Triages []Triage `gorm:"ForeignKey:IPD_ID"`
}

type Patient struct {
	gorm.Model
	//Patient_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Patient_Name string

	// Gender_ID   ทำหน้าที่เป็น FK
	Gender_ID 	*uint
	Gender    	Gender 	`gorm:"references:id"`
	Triages 	[]Triage 	`gorm:"ForeignKey:Patient_ID"`
}

type Gender struct {
	gorm.Model
	//Gender_ID		uint		`gorm:"primaryKey;autoIncrement:true"`
	Gender_Type string
	Patients []Patient `gorm:"ForeignKey:Gender_ID"`
}

type Triage struct {
	gorm.Model

	// Patient_ID ทำหน้าที่เป็น FK
	Patient_ID *uint
	Patient    Patient `gorm:"references:id"`


	// Disease_ID ทำหน้าที่เป็น FK
	Disease_ID *uint
	Disease    Disease `gorm:"references:id"`


	// IPD_ID ทำหน้าที่เป็น FK
	IPD_ID         *uint
	Ipd            Ipd `gorm:"references:id"`

	Triage_Comment string

	

	Map_Beds []Map_Bed	`gorm:"ForeignKey:Triage_ID"`
	
}

// =================================== ดึงมาจากเพื่อน ===================================

