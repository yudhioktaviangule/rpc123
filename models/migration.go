package models

import "tele-dialog-hooks/config"

func Migrate() {
	db, conn := config.DBConnect()

	db.AutoMigrate(
		&Kategori{},
		&Pelanggan{},
		&Produk{},
		&Transaksi{},
	)

	defer conn.Close()
}
