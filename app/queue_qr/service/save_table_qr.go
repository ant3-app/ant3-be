package service

import (
	"ant3/app/queue_qr/models"
	"ant3/app/queue_qr/repository"
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveTableQr(request SaveTableQrRequest) (*SaveTableQrResponse, error) ***REMOVED***
	qrTableId := primitive.NewObjectID()
	qrTableWebLink := fmt.Sprintf("%s/table-info/%s", viper.Get("ANT3_WEB_URL"), qrTableId.Hex())
	
	png, err := qrcode.Encode(qrTableWebLink, qrcode.High, 256)
	if (err != nil) ***REMOVED***
		fmt.Println("[saveQrTable] error when encoding the url link to qr code")
		return nil, err
	***REMOVED***
	
	fmt.Printf("request: %#v\n", request)
	var fileName = fmt.Sprintf("%s_table_qr_code_%s", request.MerchantId, request.TableName)
	
	fileId, err := repository.SaveImageToGDrive(png, fileName)
	if (err != nil) ***REMOVED***
		fmt.Println(map[string]interface***REMOVED******REMOVED******REMOVED***
			"png": png,
			"fileName": fileName,
		***REMOVED***,"[saveQrTable] error when saving image to gdrive")
		return nil, err
	***REMOVED***
	
	merchatId, err := primitive.ObjectIDFromHex(request.MerchantId)
	if (err != nil) ***REMOVED***
		fmt.Println(request.MerchantId, "[saveQrTable] error format string to object id")
		return nil, err
	***REMOVED***
	
	var queueQr *models.TableQr = &models.TableQr***REMOVED***
		MerchantId: merchatId,
		Name: request.TableName,
		FileId: fileId,
		Id: qrTableId,
	***REMOVED***
	
	fmt.Printf("trying to save with %#v\n", queueQr)
	id, err := repository.Save(queueQr)
	
	return &SaveTableQrResponse***REMOVED***
		Id: *id,
		FileId: fileId,
		MerchantId: merchatId.Hex(),
	***REMOVED***, nil
***REMOVED***