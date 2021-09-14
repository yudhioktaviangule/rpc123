package bussineslogic

import (
	"encoding/json"
	"tele-dialog-hooks/helpers"
	"tele-dialog-hooks/models"
)

func KategoriResponseToDialog() string {
	type pr struct {
		Text    string            `json:"text"`
		Data    []models.Kategori `json:"data"`
		Actions string            `json:"aksi"`
	}
	kat := models.KategoriGetList()
	var response pr = pr{
		Text:    "Silahkan Pilih Merk HP yang anda inginkan",
		Actions: "kat.list",
		Data:    kat,
	}

	jsonByte, err := json.Marshal(response)
	helpers.FailOnError(err, "ERROR PARSING JSON")

	return string(jsonByte)
}
