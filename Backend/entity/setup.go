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
  database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{
    Logger: SqlLogger{},
  })
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Map_Bed{},
      	&Bed{},
    		&Zone{},
		&Gender{},
		&Patient{},
		&DiseaseType{},
		&Disease{},
		&InpantientDepartment{},
		&Triage{},
		&User{},

	)
  db = database
//   // zone 1
//   db.Model(&Zone{}).Create(&Zone{
// 	Zone_Name : "หอผู้ป่วยสามัญหญิง",
//   })
//   // zone 2
//   db.Model(&Zone{}).Create(&Zone{
// 	Zone_Name : "หอผู้ป่วยสามัญชาย",
//   })
//   // zone 3
//   db.Model(&Zone{}).Create(&Zone{
// 	Zone_Name : "หอผู้ป่วยคลอด",
//   })
//    // zone 4
//    db.Model(&Zone{}).Create(&Zone{
// 	Zone_Name : "หอผู้ป่วยหลังคลอด",
//   })
//   // zone 5
//   db.Model(&Zone{}).Create(&Zone{
// 	Zone_Name : "หอผู้ป่วยพิเศษICU",
//   })
}