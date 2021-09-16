package models

import "tele-dialog-hooks/config"

type AkunBank struct {
	ID            int    `json:"id"`
	BankName      string `json:"bank_name"`
	AccountInfo   string `json:"atas_nama"`
	AccountNumber string `json:"no_rek"`
}

func AkunBankGet() (string, string, string) {
	db, sql := config.DBConnect()
	defer sql.Close()
	data := AkunBank{}
	db.First(&data)
	return data.BankName, data.AccountNumber, data.AccountInfo
}
