package controller

import (
	"ProjectSa-arm/Backend/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST gender
func CreateGender(c *gin.Context){

	var gender entity.Gender
	if err := c.ShouldBindJSON(&gender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gender})
}

// GET /getgender/:id
func GetGender(c *gin.Context) {
	var Getgender entity.Gender
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM genders WHERE id = ?", id).Scan(&Getgender).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": Getgender})
}

// GET /GetListGenders
func GetListGenders(c *gin.Context) {
	var getgenders []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM genders").Scan(&getgenders).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": getgenders})
}