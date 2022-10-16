package controller

import (
	"github.com/panupongKanin/ProjectSA-arm/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST Zone

func 	CreateZone(c *gin.Context){

	var zone entity.Zone
	if err := c.ShouldBindJSON(&zone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&zone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": zone})

}

// GET /zone/:id
func GetZone(c *gin.Context) {
	var zone entity.Zone
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM zones WHERE id = ?", id).Scan(&zone).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": zone})
}

// GET /zones
func ListZones(c *gin.Context) {
	var zones []entity.Zone
	if err := entity.DB().Raw("SELECT * FROM zones").Scan(&zones).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": zones})
}

// PATCH /zone
func UpdateZone(c *gin.Context) {
	var zone entity.Zone
	if err := c.ShouldBindJSON(&zone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", zone.ID).First(&zone); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "zone not found"})
		return
	}

	if err := entity.DB().Save(&zone).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": zone})
}
