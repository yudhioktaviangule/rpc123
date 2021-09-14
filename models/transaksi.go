package models

import (
	"fmt"
	"strings"
	"tele-dialog-hooks/config"

	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	ID              int       `gorm:"primaryKey:autoIncrement" json:"id"`
	Nomor           string    `gorm:"unique" json:"nomor"`
	Qty             int       `json:"qty"`
	HargaTransaksi  float64   `json:"harga_transaksi"`
	ProdukId        int       `json:"produk_id"`
	PelangganId     int       `json:"pelanggan_id"`
	StatusTransaksi string    `json:"status_trans" gorm:"status"`
	Pelanggan       Pelanggan `gorm:"foreignKey:pelanggan_id;references:id;constraint:onDelete:CASCADE"`
	Produk          Produk    `gorm:"foreignKey:produk_id;references:id;constraint:onDelete:CASCADE"`
	gorm.Model
}

func GetYearMonthDate() (int, string, int) {
	years := time.Now().Year()
	month := time.Now().Month().String()[0:3]
	day := time.Now().Day()
	month = strings.ToUpper(month)
	//fmt.Println(years, month, day)

	return years, month, day
}

func CreateNotrans() (string, int) {
	var count int64
	var trx Transaksi
	var isWillUse bool
	db, sql := config.DBConnect()

	db.Model(&Transaksi{}).Count(&count)
	count += 1
	y, _, d := GetYearMonthDate()
	nomor := fmt.Sprintf("%d%d-%04d", y, d, count)
	isWillUse = false
	for !isWillUse {
		nomor := fmt.Sprintf("%d%d%04d", y, d, count)
		fmt.Println(nomor)
		err := db.Where("nomor=?", nomor).First(&trx).Error
		fmt.Println(err)
		if err != nil {
			isWillUse = true
		} else {
			count -= 1
		}
	}
	defer sql.Close()
	return nomor, 1
}

func TransaksiSaveTmp(mdl Transaksi) {
	db, sql := config.DBConnect()
	mdl.StatusTransaksi = "n"
	db.Create(&mdl)
	defer sql.Close()
}

func GetTransaksiByNo(nomor string) Transaksi {
	db, sql := config.DBConnect()
	var trx Transaksi
	db.Where("nomor=?", nomor).First(&trx)
	defer sql.Close()
	return trx
}

func UpdateDataTransaksi(trx Transaksi) Transaksi {
	var tex Transaksi = trx
	db, sql := config.DBConnect()
	defer sql.Close()
	tex.StatusTransaksi = "n"
	db.Where("id=?", trx.ID).Updates(&tex)
	return tex
}
