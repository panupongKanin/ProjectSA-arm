package main

import (
	"github.com/gin-gonic/gin"
	"github.com/panupongKanin/sa-65-example/entity" 
)

func main() {
	entity.SetupDatabase()
 
  	r := gin.Default()
	
}


