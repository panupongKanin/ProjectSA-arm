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

	)
  db = database
}