package service

import (
	"ant3/app/queue_qr/models"
	"ant3/app/queue_qr/repository"
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveTableQr(request SaveTableQrRequest) (*SaveTableQrResponse, error) {
	qrTableId := primitive.NewObjectID()
	qrTableWebLink := fmt.Sprintf("%s/table-info/%s", viper.Get("ANT3_WEB_URL"), qrTableId.Hex())
	
	png, err := qrcode.Encode(qrTableWebLink, qrcode.High, 256)
	if (err != nil) {
		fmt.Println("[saveQrTable] error when encoding the url link to qr code")
		return nil, err
	}
	
	fmt.Printf("request: %#v\n", request)
	var fileName = fmt.Sprintf("%s_table_qr_code_%s", request.MerchantId, request.TableName)
	
	fileId, err := repository.SaveImageToGDrive(png, fileName)
	if (err != nil) {
		fmt.Println(map[string]interface{}{
			"png": png,
			"fileName": fileName,
		},"[saveQrTable] error when saving image to gdrive")
		return nil, err
	}
	
	merchatId, err := primitive.ObjectIDFromHex(request.MerchantId)
	if (err != nil) {
		fmt.Println(request.MerchantId, "[saveQrTable] error format string to object id")
		return nil, err
	}
	
	var queueQr *models.TableQr = &models.TableQr{
		MerchantId: merchatId,
		Name: request.TableName,
		FileId: fileId,
		Id: qrTableId,
	}
	
	fmt.Printf("trying to save with %#v\n", queueQr)
	id, err := repository.Save(queueQr)
	
	return &SaveTableQrResponse{
		Id: *id,
		FileId: fileId,
		MerchantId: merchatId.Hex(),
	}, nil
}