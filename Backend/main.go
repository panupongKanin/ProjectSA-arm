package main

import (
	"ProjectSa-arm/Backend/controller"
	"ProjectSa-arm/Backend/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
  	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST("/CreateUse",controller.CreateUser)
	r.POST("/CreateUserType",controller.CreateUserType)
	r.GET("/ListUserTypes",controller.ListUserTypes)

	r.POST("/CreateMapBed", controller.CreateMapBed)
	r.GET("/Mapbeds", controller.ListMapBeds)
	r.GET("/Mapbed/:id", controller.GetMapBed)
	r.Run()
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