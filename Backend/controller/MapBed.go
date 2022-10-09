package controller

import (
	"ProjectSa-arm/Backend/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST Map_Bed

func 	CreateMapBed(c *gin.Context){

	var triage entity.Triage
	// var bed entity.Bed
	// var user entity.User

	var map_bed entity.Map_Bed

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&map_bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา triage ด้วย id
	if tx := entity.DB().Where("id = ?", map_bed.Triage_ID).First(&triage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

}

// GET /Mapbed/:id
func GetMapBed(c *gin.Context) {
	var map_bed entity.Map_Bed
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM map_beds WHERE id = ?", id).Scan(&map_bed).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": map_bed})
}

// GET /mapbeds
func ListMapBeds(c *gin.Context) {
	var map_beds []entity.Map_Bed
	if err := entity.DB().Raw("SELECT * FROM map_beds").Scan(&map_beds).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": map_beds})
}


