package main

import (
	"ProjectSa-arm/Backend/controller"
	"ProjectSa-arm/Backend/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
  	r := gin.Default()
	r.POST("/CreateUse",controller.CreateUser)
	r.POST("/CreateUserType",controller.CreateUserType)
	r.GET("/ListUserTypes",controller.ListUserTypes)

	r.POST("/CreateMapBed", controller.CreateMapBed)
	r.GET("/Mapbeds", controller.ListMapBeds)
	r.GET("/Mapbed/:id", controller.GetMapBed)
	r.Run()
}


// func CORMiddleware() gin.HandlerFunc{

// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set(){set}
// 	}
// }