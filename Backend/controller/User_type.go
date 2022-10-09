package controller

import (
	"ProjectSa-arm/Backend/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST User_Type

func 	CreateUserType(c *gin.Context){

	var user_type entity.User_Type
	if err := c.ShouldBindJSON(&user_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&user_type).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user_type})

}

// GET /user_type/:id
func GetUserType(c *gin.Context) {
	var user_type entity.User_Type
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM user_types WHERE id = ?", id).Scan(&user_type).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": user_type})
}

// GET /user_types
func ListUserTypes(c *gin.Context) {
	var user_types []entity.User_Type
	if err := entity.DB().Raw("SELECT * FROM user_types").Scan(&user_types).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	c.JSON(http.StatusOK, gin.H{"data": user_types})
}