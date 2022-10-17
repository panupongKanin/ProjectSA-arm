package controller

import (
	"github.com/panupongKanin/ProjectSA-arm/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST Patient
func CreatePatient(c *gin.Context){

	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET /getPatient/:id
func GetPatient(c *gin.Context) {
	var GetPatient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&GetPatient).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": GetPatient})
}

// GET /GetListPatients
func GetListPatients(c *gin.Context) {
	var getPatients []entity.Patient
	if err := entity.DB().Preload("Gender").Raw("SELECT * FROM patients").Find(&getPatients).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": getPatients})
}