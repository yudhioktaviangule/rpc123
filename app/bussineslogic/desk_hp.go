package bussineslogic

import (
	"encoding/json"
	"fmt"
	"strings"
	"tele-dialog-hooks/helpers"
	"tele-dialog-hooks/models"
)

func ProductDesc(textChat string) string {
	type resp struct {
		Text    string        `json:"text"`
		Produk  models.Produk `json:"data"`
		Actions string        `json:"aksi"`
	}

	var response resp
	splitter := strings.Split(strings.ToLower(textChat), " deskripsi ")
	name := splitter[1]
	hp := models.ProductGetByName(name)
	response.Actions = "prd.desc"
	response.Text = fmt.Sprintf("Spesifikasi handphone %s", name)

	response.Produk = hp
	marshal, err := json.Marshal(response)
	helpers.FailOnError(err, "ERROR PARSE")
	return string(marshal)
}
