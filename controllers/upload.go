package controllers

import (
	"net/http"
	"pr/models"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
        return
    }

    // Validasi ukuran dan tipe file
    if file.Size > 2*1024*1024 || (file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png") {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type or size"})
        return
    }

    err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
        return
    }

    // Simpan nama file ke dalam database (misalnya setelah laporan dibuat)
    var laporan models.LaporanPertandingan
    if err := models.DB.Where("id = ?", c.Param("id")).First(&laporan).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Laporan not found"})
        return
    }
    laporan.FotoPertandingan = file.Filename
    models.DB.Save(&laporan)

    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
