package controller

import (
	"ProjectSa-arm/Backend/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST Triage
func CreateTriage(c *gin.Context){

	var triage entity.Triage
	if err := c.ShouldBindJSON(&triage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&triage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": triage})
}

// GET /getPatient/:id
func GetTriage(c *gin.Context) {
	var GetTriage entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM triages WHERE id = ?", id).Scan(&GetTriage).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetTriage})
}

// GET /GetListPatients
func GetListTriages(c *gin.Context) {
	var getListTriages []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM triages").Scan(&getListTriages).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": getListTriages})
}