package controllers

import (
	"net/http"
	"pr/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var laporan []models.LaporanPertandingan

	models.DB.Find(&laporan)
	c.JSON(http.StatusOK, gin.H{
		"data": laporan,
	})
}

func Show(c *gin.Context) {
    var laporan models.LaporanPertandingan
    if err := models.DB.Where("id = ?", c.Param("id")).First(&laporan).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Laporan not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": laporan})
}


func Create(c *gin.Context) {
	var laporan models.LaporanPertandingan

	if err := c.ShouldBindJSON(&laporan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Create(&laporan)
	c.JSON(http.StatusOK, gin.H{
		"data": laporan,
	})
}

func Update(c *gin.Context) {
    var laporan models.LaporanPertandingan
    if err := models.DB.Where("id = ?", c.Param("id")).First(&laporan).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Laporan not found"})
        return
    }

    if err := c.ShouldBindJSON(&laporan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Save(&laporan)
    c.JSON(http.StatusOK, gin.H{"data": laporan})
}


func Delete(c *gin.Context) {
    var laporan models.LaporanPertandingan
    if err := models.DB.Where("id = ?", c.Param("id")).First(&laporan).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Laporan not found"})
        return
    }

    models.DB.Delete(&laporan)
    c.JSON(http.StatusOK, gin.H{"data": "Laporan deleted successfully"})
}

