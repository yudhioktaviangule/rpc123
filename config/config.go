package config

import (
	"database/sql"
	"fmt"
	"tele-dialog-hooks/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB_USER, DB_HOST, DB_PASSWORD, DB_PORT, DB_NAME string
var NOMOR_REK, BANK, ATAS_NAMA string

func InitEnv() {
	DB_USER = "yudhi"
	DB_PASSWORD = "123"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_NAME = "db_penjualan"
	BANK = "BCA"
	NOMOR_REK = "871221092"
	ATAS_NAMA = "YUDHI"
}

func DBConnect() (*gorm.DB, *sql.DB) {
	tcp := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(tcp), &gorm.Config{})
	helpers.FailOnError(err, "Fail Connection Database")
	sql, _ := db.DB()
	return db, sql
}
