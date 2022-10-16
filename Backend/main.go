package main
 
import (
  "github.com/panupongKanin/ProjectSA-arm/controller"
  "github.com/panupongKanin/ProjectSA-arm/entity"
  "github.com/gin-gonic/gin"
)

func main() {
  	entity.SetupDatabase()


  	r := gin.Default()
	r.Use(CORSMiddleware())
	//=========== Main Table Mapping Bed ===========
	r.POST("/CreateMapBed",controller.CreateMapBed)
	r.GET("/GetListMapBeds",controller.GetListMapBeds)
	r.GET("/GetMapBed/:id",controller.GetMapBed)
	//=========== Main Table Mapping Bed ===========


	r.POST("/CreateUserType",controller.CreateUser)
	r.GET("/ListUserTypes",controller.ListUsers)

	//===========Zone===========
	r.POST("/CreateZone",controller.CreateZone)
	r.GET("/GetListZones",controller.ListZones)
	//===========Bed===========
	r.POST("/CreateBed",controller.CreateBed)
	r.GET("/GetListBeds",controller.ListBeds)
	r.GET("/Bed/:zoneid",controller.GetBed_by_zone)

	//===========gender===========
	r.POST("/CreateGender",controller.CreateGender)
	r.GET("/GetListGenders",controller.GetListGenders)
	r.GET("/GetGender/:id",controller.GetGender)

	//===========IPD===========
	r.POST("/CreateIPD",controller.CreateIPD)
	r.GET("/GetListIPDs",controller.GetListIPDs)
	r.GET("/GetIPD/:id",controller.GetIPD)

	//===========Patient===========
	r.POST("/CreatePatient",controller.CreatePatient)
	r.GET("/GetListPatients",controller.GetListPatients)
	r.GET("/GetPatient/:id",controller.GetPatient)

	//===========DiseaseType===========
	r.POST("/CreateDiseaseType",controller.CreateDiseaseType)
	r.GET("/GetListDiseaseType",controller.GetListDiseaseTypes)
	r.GET("/GetDiseaseType/:id",controller.GetDiseaseType)

	//===========Disease===========
	r.POST("/CreateDisease",controller.CreateDisease)
	r.GET("/GetListDisease",controller.GetListDiseases)
	r.GET("/GetDisease/:id",controller.GetDisease)

	//===========Triage===========
	r.POST("/CreateTriage",controller.CreateTriage)
	r.GET("/GetListTriages",controller.GetListTriages)
	r.GET("/GetTriage/:id",controller.GetTriage)
 
	

	//===========Test===========

	r.GET("/Gettest",controller.Gettest)



  	r.Run("localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
     
		if c.Request.Method == "OPTIONS" {
		  c.AbortWithStatus(204)
		  return
		}
     
		c.Next()
	}
}