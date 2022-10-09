package entity
 
import (
           "gorm.io/gorm"
           "gorm.io/driver/sqlite"
           "time"
)
 
var db *gorm.DB
 
func DB() *gorm.DB {
           return db
}
 
func SetupDatabase() {
  database, err := gorm.Open(sqlite.Open("Map_Bed.db"), &gorm.Config{})
  if err != nil {
        panic("failed to connect database")
  }
 
  // Migrate the schema
  database.AutoMigrate(
    &User{},
    &User_Type{},
    &Zone{},
    &Bed{},
    &Map_Bed{},

    &Triage{},
    &Patient{},
    &Gender{},
    &Disease{},
    &Disease_Type{},
    &Ipd{},
    )
 
  db = database


  var mapbed = Map_Bed{Triage_ID: "1", Admidtime: time.Now(),Bed_ID:"1",Map_Bed_Comment:"-",User_ID:"1"}
  db.Create(&mapbed)
}

