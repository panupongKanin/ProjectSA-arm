package controller

// import (
// 	"ProjectSa-arm/Backend/entity"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// // POST Zone

// func 	CreateZone(c *gin.Context){

// 	var zone entity.Zone
// 	if err := c.ShouldBindJSON(&zone); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := entity.DB().Create(&zone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": zone})

// }

// // GET /zone/:id
// func GetZone(c *gin.Context) {
// 	var zone entity.Zone
// 	id := c.Param("id")
// 	if err := entity.DB().Raw("SELECT * FROM zones WHERE id = ?", id).Scan(&zone).Error; err != nil {
// 		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		 return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": zone})
// }

// // GET /zones
// func ListZones(c *gin.Context) {
// 	var zones []entity.Map_Bed
// 	if err := entity.DB().Raw("SELECT * FROM zones").Scan(&zones).Error; err != nil {
// 		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		 return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": zones})
// }
