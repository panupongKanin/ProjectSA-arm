package controller

import (
	"github.com/panupongKanin/ProjectSA-arm/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// POST Map_Bed

func 	CreateMapBed(c *gin.Context){

	var triage entity.Triage
	var bed entity.Bed
	//var user entity.User

	var map_bed entity.Map_Bed

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 10 จะถูก bind เข้าตัวแปร map_bed
	if err := c.ShouldBindJSON(&map_bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา triage ด้วย id
	if tx := entity.DB().Where("id = ?", map_bed.Triage_ID).First(&triage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "triage not found"})
		return
	}

	// 12: ค้นหา bed ด้วย id
	if tx := entity.DB().Where("id = ?", map_bed.Bed_ID).First(&bed); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bed not found"})
		return
	}

	// 13: ค้นหา user ด้วย id
	// if tx := entity.DB().Where("id = ?", map_bed.User_ID).First(&user); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
	// 	return
	// }

	// 14: สร้าง Mapbed
	mb := entity.Map_Bed{
		Triage_ID: map_bed.Triage_ID,         	// โยงความสัมพันธ์กับ Entity Triage
		Admidtime: map_bed.Admidtime, // ตั้งค่าฟิลด์ Admidtime
		Bed_ID:    map_bed.Bed_ID,                 // โยงความสัมพันธ์กับ Entity Bed
		MapBed_Comment: map_bed.MapBed_Comment,
		User_ID: map_bed.User_ID,               	// โยงความสัมพันธ์กับ Entity User
	}

	// 15: บันทึก
	if err := entity.DB().Create(&mb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mb})

}

// GET /Mapbed/:id
func GetMapBed(c *gin.Context) {
	var GetMapBed entity.Map_Bed
	id := c.Param("id")
	if err := entity.DB().Preload("Bed.Zone").Preload("Triage.Patient.Gender").Preload("Triage.Disease.Disease_Type").Preload("Triage.Ipd").Preload("User.User_Type").Raw("SELECT * FROM map_beds WHERE id = ?", id).Scan(&GetMapBed).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": GetMapBed})
}

// GET /mapbeds
func GetListMapBeds(c *gin.Context) {
	var GetMapBeds []entity.Map_Bed
	if err := entity.DB().Preload("Bed.Zone").Preload("Triage.Patient.Gender").Preload("Triage.Disease.Disease_Type").Preload("Triage.Ipd").Preload("User.User_Type").Raw("SELECT * FROM map_beds").Find(&GetMapBeds).Error; err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	c.JSON(http.StatusOK, gin.H{"data": GetMapBeds})
}


