package controller

type tableQrRequest struct {
	MerchantId string `json:"merchantId"`
	TableName string `json:"tableName"`
}

type tableQueueRequest struct {
	TableQrId string `json:"tableQrId"`
}