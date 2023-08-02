package qpay_v1

import (
	"encoding/json"
)

type qpay struct {
	endpoint      string
	client_id     string
	client_secret string
	grant_type    string
	refresh_token string
	callback      string
	invoiceCode   string
	merchantId    string
	templateId    string
	branchId      string
	posId         string
	loginObject   *QpayLoginResponse
}

type QPay interface {
	CreateInvoice(input QPayInvoiceCreateRequest) (QPaySimpleInvoiceResponse, error)
	GetInvoice(invoiceId string) (QpayInvoiceGetResponse, error)
	CheckPayment(paymentID string) (QpayPaymentCheckResponse, error)
}

func New(client_id, client_secret, endpoint, callback, merchantId, templateId, branchId, posId string) QPay {
	return &qpay{
		endpoint:      endpoint,
		client_id:     client_id,
		client_secret: client_secret,
		grant_type:    "client",
		refresh_token: "",
		callback:      callback,
		merchantId:    merchantId,
		templateId:    templateId,
		branchId:      branchId,
		posId:         posId,
		loginObject:   nil,
	}
}

func (q *qpay) CreateInvoice(input QPayInvoiceCreateRequest) (QPaySimpleInvoiceResponse, error) {
	input.BranchID = q.branchId
	input.PosID = q.posId
	input.MerchantID = q.merchantId
	input.TemplateID = q.templateId

	res, err := q.httpRequestQPay(input, QPayInvoiceCreate, "")
	if err != nil {
		return QPaySimpleInvoiceResponse{}, err
	}

	var response QPaySimpleInvoiceResponse
	json.Unmarshal(res, &response)

	return response, nil
}
func (q *qpay) GetInvoice(invoiceId string) (QpayInvoiceGetResponse, error) {
	res, err := q.httpRequestQPay(nil, QPayInvoiceGet, invoiceId)
	if err != nil {
		return QpayInvoiceGetResponse{}, err
	}

	var response QpayInvoiceGetResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpay) CheckPayment(paymentID string) (QpayPaymentCheckResponse, error) {
	var response QpayPaymentCheckResponse

	res, err := q.httpRequestQPay(nil, QPayPaymentCheck, paymentID)
	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}
