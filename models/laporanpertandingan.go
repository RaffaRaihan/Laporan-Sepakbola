package models

import "gorm.io/gorm"

type LaporanPertandingan struct {
    gorm.Model
    NamaTim         string  `json:"nama_tim" binding:"required"`
    SkorAkhir       int     `json:"skor_akhir" binding:"required"`
    Tanggal         string  `json:"tanggal" binding:"required"`
    PencetakGol     string  `json:"pencetak_gol" binding:"required"`
    CatatanTambahan string  `json:"catatan_tambahan"`
    FotoPertandingan string `json:"foto_pertandingan"` // Nama file foto yang di-upload
}
