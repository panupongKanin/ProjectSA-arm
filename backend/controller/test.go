package controller

import (

	"fmt"
	"log"
	"net/http"
	//"strings"

	"github.com/gin-gonic/gin"
	"github.com/panupongKanin/ProjectSA-arm/entity"
)



func Gettest(c *gin.Context) {
	
	


	rows, err := entity.DB().Table("zones").Select("zones.zone_name, beds.bed_name").Joins("inner join beds on zones.id = beds.zone_id").Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tests := make([]string,0)
	// var bed_name string
	var zone_name	string
	var Bed_Name	string
	for rows.Next(){

		
		if err := rows.Scan(&zone_name,&Bed_Name);err !=  nil {
			log.Fatal(err)
		}
		tests = append(tests,zone_name,Bed_Name)
		fmt.Printf("%v\n", tests)
		
		
		
	

	
	}
	c.JSON(http.StatusOK, gin.H{"data": tests})
	// t:=strings.Join(tests,"")

	fmt.Printf("%v\n", tests)
	// fmt.Printf("%v\n", t)

	



}
// rows, err := db.Table("zones").Select("zones.zone_name, beds.bed_name").Joins("inner join beds on zones.id = beds.zone_id").Rows()
// rows,err:= entity.DB().Table("patients").Select("patients.patient_name").Joins("inner join triages on patients.id = triages.patient_id").Rows()