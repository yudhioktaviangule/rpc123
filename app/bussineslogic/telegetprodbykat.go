package bussineslogic

import (
	"encoding/json"
	"fmt"
	"strings"
	"tele-dialog-hooks/helpers"
	"tele-dialog-hooks/models"
)

func KatGetName(text string) string {
	spliits := strings.Split(text, " ")
	texie := ""
	for _, split := range spliits {
		texie = split
	}

	return texie
}

func ProductListResponse(products []models.Produk, katname string) string {
	type pr struct {
		Text    string          `json:"text"`
		Produks []models.Produk `json:"list_handphone"`
		Actions string          `json:"aksi"`
	}
	var arrs pr
	for _, produk := range products {
		arrs.Produks = append(arrs.Produks, produk)
	}
	arrs.Actions = "prd.list"
	arrs.Text = fmt.Sprintf("Daftar Handphone dari Merk %s", katname)
	jsonByte, err := json.Marshal(arrs)
	helpers.FailOnError(err, "ERROR PARSING JSON")

	return string(jsonByte)
}
