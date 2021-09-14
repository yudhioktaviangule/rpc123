package models

import (
	"tele-dialog-hooks/config"

	"gorm.io/gorm"
)

type Produk struct {
	ID         int      `gorm:"primaryKey:autoIncrement" json:"id"`
	Inisial    string   `json:"inisial" gorm:"unique"`
	Name       string   `json:"name"`
	Harga      float64  `json:"harga"`
	Gambar     string   `json:"gambar"`
	Ram        string   `json:"ram"`
	Memori     string   `json:"memori"`
	Stok       int      `json:"stok"`
	KategoriId int      `json:"kategori_id"`
	Kategori   Kategori `gorm:"foreignKey:kategori_id;references:id;constraint:onDelete:CASCADE"`
	gorm.Model
}

func ProductGetByKat(kat_id int) []Produk {
	db, sql := config.DBConnect()
	defer sql.Close()
	var results []Produk
	db.Where("kategori_id=?", kat_id).Find(&results)
	return results

}

func ProductGetByName(name string) Produk {
	db, sql := config.DBConnect()
	defer sql.Close()
	var result Produk
	db.Where("name=?", name).First(&result)

	return result
}

func ProductList() []Produk {
	db, sql := config.DBConnect()
	defer sql.Close()
	var result []Produk
	db.Find(&result)
	return result
}
