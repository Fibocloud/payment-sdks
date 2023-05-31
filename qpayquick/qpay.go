package qpayquick

import (
	"encoding/json"
)

type qpayquick struct {
	endpoint    string
	password    string
	username    string
	callback    string
	invoiceCode string
	merchantId  string
	terminalID  string
	loginObject *qpayLoginResponse
}

type QPayQuick interface {
	CreateCompany(input QpayCompanyCreateRequest) (QpayCompanyCreateResponse, error)
	CreatePerson(input QpayPersonCreateRequest) (QpayPersonCreateResponse, error)
	GetMerchant(merchantID string) (QpayMerchantGetResponse, error)
	ListMerchant(input QpayOffset) (QpayMerchantListResponse, error)
	CreateInvoice(input QpayInvoiceRequest) (QpayInvoiceResponse, error)
	GetInvoice(invoiceId string) (QpayInvoiceGetResponse, error)
	CheckPayment(invoiceID string) (QpayPaymentCheckResponse, error)
}

func New(username, password, endpoint, callback, invoiceCode, merchantId string) QPayQuick {
	return &qpayquick{
		endpoint:    endpoint,
		password:    password,
		username:    username,
		callback:    callback,
		invoiceCode: invoiceCode,
		merchantId:  merchantId,
		loginObject: nil,
	}
}

func (q *qpayquick) CreateCompany(input QpayCompanyCreateRequest) (QpayCompanyCreateResponse, error) {

	res, err := q.httpRequestQPay(input, QPayCreateCompany, "")
	if err != nil {
		return QpayCompanyCreateResponse{}, err
	}

	var response QpayCompanyCreateResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpayquick) CreatePerson(input QpayPersonCreateRequest) (QpayPersonCreateResponse, error) {
	res, err := q.httpRequestQPay(input, QPayCreatePerson, "")
	if err != nil {
		return QpayPersonCreateResponse{}, err
	}

	var response QpayPersonCreateResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpayquick) GetMerchant(merchantID string) (QpayMerchantGetResponse, error) {
	res, err := q.httpRequestQPay(nil, QPayGetMerchant, merchantID)
	if err != nil {
		return QpayMerchantGetResponse{}, err
	}

	var response QpayMerchantGetResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpayquick) ListMerchant(input QpayOffset) (QpayMerchantListResponse, error) {
	request := QpayMerchantListRequest{
		Offset: input,
	}
	res, err := q.httpRequestQPay(request, QPayMerchantList, "")
	if err != nil {
		return QpayMerchantListResponse{}, err
	}

	var response QpayMerchantListResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpayquick) CreateInvoice(input QpayInvoiceRequest) (QpayInvoiceResponse, error) {
	res, err := q.httpRequestQPay(input, QPayInvoiceCreate, "")
	if err != nil {
		return QpayInvoiceResponse{}, err
	}
	var response QpayInvoiceResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpayquick) GetInvoice(invoiceId string) (QpayInvoiceGetResponse, error) {
	res, err := q.httpRequestQPay(nil, QPayInvoiceGet, invoiceId)
	if err != nil {
		return QpayInvoiceGetResponse{}, err
	}
	var response QpayInvoiceGetResponse
	json.Unmarshal(res, &response)

	return response, nil
}
func (q *qpayquick) CheckPayment(invoiceID string) (QpayPaymentCheckResponse, error) {
	request := QpayPaymentCheckRequest{
		InvoiceID: invoiceID,
	}
	res, err := q.httpRequestQPay(request, QPayPaymentCheck, "")
	if err != nil {
		return QpayPaymentCheckResponse{}, err
	}

	var response QpayPaymentCheckResponse
	json.Unmarshal(res, &response)

	return response, nil
}
