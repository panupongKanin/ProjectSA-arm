package controller

// import (
// 	"ProjectSa-arm/Backend/entity"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// // POST Bed
// func CreateBed(c *gin.Context){

// 	var bed entity.Bed
// 	if err := c.ShouldBindJSON(&bed); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := entity.DB().Create(&bed).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": bed})
// }

// // GET /bed/:id
// func GetBed(c *gin.Context) {
// 	var bed entity.Map_Bed
// 	id := c.Param("id")
// 	if err := entity.DB().Raw("SELECT * FROM beds WHERE id = ?", id).Scan(&bed).Error; err != nil {
// 		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		 return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": bed})
// }

// // GET /beds
// func ListBeds(c *gin.Context) {
// 	var beds []entity.Map_Bed
// 	if err := entity.DB().Raw("SELECT * FROM beds").Scan(&beds).Error; err != nil {
// 		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		 return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": beds})
// }