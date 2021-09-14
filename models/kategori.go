package models

import (
	"tele-dialog-hooks/config"

	"gorm.io/gorm"
)

type Kategori struct {
	ID   int    `gorm:"primaryKey:autoIncrement" json:"id"`
	Name string `json:"name"`
	gorm.Model
}

func KategoriGetList() []Kategori {
	db, sql := config.DBConnect()
	var result []Kategori
	db.Find(&result)
	defer sql.Close()
	return result

}

func KatGet(name string) Kategori {
	db, sql := config.DBConnect()
	defer sql.Close()
	var result Kategori
	db.Where("name=?", name).First(&result)
	return result
}
