package main

import (
	"ProjectSa-arm/Backend/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	
	entity.SetupDatabase()
  	r := gin.Default()
	r.Run()
	
}


