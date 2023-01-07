package service

type SaveTableQrRequest struct {
	MerchantId string
	TableName string
}

type SaveTableQrResponse struct {
	FileId string
	MerchantId string
	Id string
}