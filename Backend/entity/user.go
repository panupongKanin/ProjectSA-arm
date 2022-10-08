package entity

import (

	"gorm.io/gorm"
)

type user struct {
  gorm.Model
  UserID          uint          `gorm:"primaryKey;autoIncrement:true"`
  UserName        string        `gorm:"uniqueIndex"`
  UserPassword    string
  UserType        userType      `gorm:"references:Usertype_ID"`
  Name            string
  
}
