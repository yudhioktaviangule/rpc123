package controllers

import (
	"tele-dialog-hooks/app/bussineslogic"
	_structs "tele-dialog-hooks/controllers/structs"
	"tele-dialog-hooks/helpers"
	"tele-dialog-hooks/models"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	pb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func TeleWebhooks(ctx *gin.Context) {
	var telmes _structs.TeleMessageRequest
	ctx.ShouldBindJSON(&telmes)
	disp := telmes.QueryResult.Intent.DisplayName

	texo := ""
	switch disp {
	case "pencarian_kategori":

		texo = bussineslogic.KategoriResponseToDialog()
	case "list_by_kat":
		kat_name := bussineslogic.KatGetName(telmes.QueryResult.QueryText)
		kategori := models.KatGet(kat_name)
		products := models.ProductGetByKat(kategori.ID)
		texo = bussineslogic.ProductListResponse(products, kategori.Name)
	case "deskripsi_handphone":
		texo = bussineslogic.ProductDesc(telmes.QueryResult.QueryText)
	case "jadi_beli":
		texo = bussineslogic.JadiBeliHape(telmes.QueryResult.QueryText)
	case "prc_transaksi":
		texo = bussineslogic.UpdateTransaksiHP(telmes.QueryResult.QueryText)
	}
	res := &pb.WebhookResponse{FulfillmentText: texo}
	m := &jsonpb.Marshaler{}
	err := m.Marshal(ctx.Writer, res)

	if err != nil {
		helpers.FailOnError(err, "ERROR")
		return
	}

}
