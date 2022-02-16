package socialpay

import "github.com/Fibocloud/payment-sdks/utils"

func mapToSimpleResponse(resp map[string]interface{}) (response SocialPaySimpleResponse) {
	response = SocialPaySimpleResponse{
		Description: utils.GetValidString(resp["desc"]),
		Status:      utils.GetValidString(resp["status"]),
	}
	return
}

func mapToTransactionInfo(resp map[string]interface{}) (response SocialPayTransactionResponse) {
	response = SocialPayTransactionResponse{
		ApprovalCode:        utils.GetValidString(resp["approval_code"]),
		Amount:              utils.GetValidFloat(resp["amount"]),
		CardNumber:          utils.GetValidString(resp["card_number"]),
		ResponseDescription: utils.GetValidString(resp["resp_desc"]),
		ResponseCode:        utils.GetValidString(resp["resp_code"]),
		Terminal:            utils.GetValidString(resp["terminal"]),
		Invoice:             utils.GetValidString(resp["invoice"]),
		Checksum:            utils.GetValidString(resp["checksum"]),
	}
	return
}

func mapToSettlementResponse(resp map[string]interface{}) (response SocialPayPaymentSettlementResponse) {
	response = SocialPayPaymentSettlementResponse{
		Amount: utils.GetValidFloat(resp["amount"]),
		Count:  resp["count"].(int),
		Status: utils.GetValidString(resp["status"]),
	}
	return
}

func mapToErrorResponse(resp map[string]interface{}) (response SocialPayErrorResponse) {
	return SocialPayErrorResponse{
		ErrorDescription: utils.GetValidString(resp["errorDesc"]),
		ErrorType:        utils.GetValidString(resp["errorType"]),
	}
}
