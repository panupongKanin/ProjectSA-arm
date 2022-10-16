package entity

import (
	"context"
	"fmt"


	"time"


	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
 
var db *gorm.DB
 
func DB() *gorm.DB {
           return db
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}
 
func SetupDatabase() {
  database, err := gorm.Open(sqlite.Open("Map_Bed.db"), &gorm.Config{
    Logger: SqlLogger{},
  })
  if err != nil {
        panic("failed to connect database")
  }
 
  // Migrate the schema
  database.AutoMigrate(
    &Zone{},
    &Bed{},

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
 



 

  


 

  
  // rows, err := db.Table("zones").Select("zones.zone_name, beds.bed_name").Joins("inner join beds on zones.id = beds.zone_id").Rows()
	// tests := make([]string,0)
	// for rows.Next(){
	// 	var zone_name string
  //   var bed_name string
	// 	rows.Scan(&zone_name,&bed_name);
	//   tests = append(tests,"{" + "\"" + "Zonename" + "\"" + ":" + "\"" + zone_name + "\"" +",","\"" + "Bedname" + "\"" + ":" + "\"" + bed_name + "\"" + "}")
		

  // }
  // fmt.Printf("%v\n", tests)
}