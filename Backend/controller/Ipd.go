package controller

import (
	"ProjectSa-arm/Backend/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST gender
func CreateIPD(c *gin.Context){

	var ipd entity.Ipd
	if err := c.ShouldBindJSON(&ipd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ipd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ipd})
}

// GET /getipd/:id
func GetIPD(c *gin.Context) {
	var GetIPD entity.Ipd
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM ipds WHERE id = ?", id).Scan(&GetIPD).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetIPD})
}

// GET /GetListIPDs
func GetListIPDs(c *gin.Context) {
	var getipds []entity.Ipd
	if err := entity.DB().Raw("SELECT * FROM ipds").Scan(&getipds).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": getipds})
}