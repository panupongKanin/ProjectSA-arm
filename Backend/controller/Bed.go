package controller

import (
	"ProjectSa-arm/Backend/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST Bed
func CreateBed(c *gin.Context){

	var bed entity.Bed
	if err := c.ShouldBindJSON(&bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&bed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bed})
}

// GET /bed/:id
func GetBed(c *gin.Context) {
	var bed entity.Bed
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM beds WHERE id = ?", id).Scan(&bed).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": bed})
}

// GET /bed/:zoneid
func GetBed_by_zone(c *gin.Context) {
	var bed_by_zone []entity.Bed
	zone_id := c.Param("zoneid")
	if err := entity.DB().Raw("SELECT * FROM beds WHERE Zone_ID = ?", zone_id).Scan(&bed_by_zone).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": bed_by_zone})
}

// GET /beds
func ListBeds(c *gin.Context) {
	var beds []entity.Bed
	if err := entity.DB().Raw("SELECT * FROM beds").Scan(&beds).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": beds})
}

// PATCH /users
func UpdateBed(c *gin.Context) {
	var bed entity.Bed
	if err := c.ShouldBindJSON(&bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bed.ID).First(&bed); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bed not found"})
		return
	}

	if err := entity.DB().Save(&bed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bed})
}