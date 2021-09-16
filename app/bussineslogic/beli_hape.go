package bussineslogic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tele-dialog-hooks/models"
)

func JadiBeliHape(queryText string) string {
	type resp struct {
		Text    string           `json:"text"`
		Data    models.Transaksi `json:"data"`
		Actions string           `json:"aksi"`
	}
	qtext := strings.ToLower(queryText)
	splits := strings.Split(qtext, "saya ingin membeli ")
	nomor, id_pelanggan_tmp := models.CreateNotrans()
	produk := models.ProductGetByName(splits[1])
	jumlah_beli := 1
	var mdl models.Transaksi
	mdl.Nomor = nomor
	mdl.PelangganId = id_pelanggan_tmp
	mdl.Qty = jumlah_beli
	mdl.ProdukId = produk.ID
	mdl.HargaTransaksi = produk.Harga
	models.TransaksiSaveTmp(mdl)
	var response resp = resp{
		Text:    fmt.Sprintf("Masukkan identitas anda dengan format transaksi<spasi>beli<spasi>Nama anda#alamat lengkap#nomor telepon#jumlah beli#%s", nomor),
		Actions: "beli.identitas",
		Data:    mdl,
	}
	jsonM, _ := json.Marshal(response)
	return string(jsonM)
}

func UpdateTransaksiHP(qtext string) string {
	type resp struct {
		Text    string           `json:"text"`
		Data    models.Transaksi `json:"data"`
		Actions string           `json:"aksi"`
	}
	var trans resp
	respon_split := strings.Split(qtext, "beli ")
	resource_split := strings.Split(respon_split[1], "#")
	var transkasiku models.Transaksi = models.GetTransaksiByNo(resource_split[4])
	tqt, err := strconv.Atoi(resource_split[3])
	if err != nil {
		tqt = 1
	}

	var pelanggan models.Pelanggan = models.Pelanggan{
		Name:    resource_split[0],
		Alamat:  resource_split[1],
		Telepon: resource_split[2],
	}
	pelanggan = models.InsertPelanggan(pelanggan)
	transkasiku.PelangganId = pelanggan.ID
	transkasiku.Qty = tqt
	transkasiku = models.UpdateDataTransaksi(transkasiku)
	trans.Actions = "beli.done"
	nama, norek, an := models.AkunBankGet()
	harga := transkasiku.HargaTransaksi * float64(transkasiku.Qty)
	trans.Text = fmt.Sprintf("Pembelian Berhasil Silahkan melakukan pembayaran melalui Bank %s ke nomor rekening %s Atas nama %s sejumlah %.0f", nama, norek, an, harga)
	trans.Data = transkasiku
	fmt.Println(transkasiku.PelangganId)
	jsonM, _ := json.Marshal(trans)
	return string(jsonM)
}
