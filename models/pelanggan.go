package models

import (
	"tele-dialog-hooks/config"

	"gorm.io/gorm"
)

type Pelanggan struct {
	ID      int    `gorm:"primaryKey:autoIncrement" json:"id"`
	Name    string `json:"name"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
	gorm.Model
}

func InsertPelanggan(pelanggan Pelanggan) Pelanggan {
	db, sql := config.DBConnect()
	var xpel Pelanggan = pelanggan
	db.Create(&xpel)
	defer sql.Close()
	return xpel
}
