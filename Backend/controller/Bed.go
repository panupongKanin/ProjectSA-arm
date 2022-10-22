package controller

import (
	"github.com/panupongKanin/ProjectSA-arm/entity"
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
func GetBedName(c *gin.Context) {
	var bed entity.Bed
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT bed_name FROM beds WHERE id = ?", id).Scan(&bed).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": bed})
}

// GET /bed/:zoneid
func GetBed_by_zone(c *gin.Context) {
	var bed_by_zone []entity.Bed
	zone_id := c.Param("zoneid")
	if err := entity.DB().Raw("SELECT * FROM beds WHERE Zone_ID = ? AND Bed_State = 0", zone_id).Scan(&bed_by_zone).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": bed_by_zone})
}

// GET /beds
func ListBeds(c *gin.Context) {
	var beds []entity.Bed
	if err := entity.DB().Preload("Zone").Raw("SELECT * FROM beds").Find(&beds).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": beds})
}

func UpdateBedstate(c *gin.Context) {
	var bed entity.Bed
	if err := c.ShouldBindJSON(&bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Model(bed).Where("id = ?", bed.ID).Update("Bed_State",bed.Bed_State).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bed})
}