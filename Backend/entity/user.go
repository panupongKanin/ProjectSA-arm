package entity

import (


	"gorm.io/gorm"
)
 
type USER struct {
  gorm.Model
  USER_ID         uint
  USER_NAME       string
  USER_PASSWORD   string `gom:"uniqueIndex"`
  USER_TYPE       string 
  NAME            string `gom:"uniqueIndex"`
}

type ZONE struct {
  gorm.Model
  ZONE_ID         uint
  ZONE_NAME       string
}

type BED struct{
  gorm.Model
  BED_ID          uint
  ZONE_ID         uint ``
}