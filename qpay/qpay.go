package qpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type qpay struct {
	endpoint    string
	password    string
	username    string
	callback    string
	invoiceCode string
	merchantId  string
	loginObject *qpayLoginResponse
}

type QPay interface {
	CreateInvoice(input QPayCreateInvoiceInput) (QPaySimpleInvoiceResponse, error)
	GetInvoice(invoiceId string) (QpayInvoiceGetResponse, error)
	CancelInvoice(invoiceId string) (interface{}, error)
	GetPayment(invoiceId string) (interface{}, error)
	CheckPayment(invoiceId string, pageLimit, pageNumber int32) (QpayPaymentCheckResponse, error)
	CancelPayment(invoiceId, paymentUUID string) (QpayPaymentCheckResponse, error)
	RefundPayment(invoiceId, paymentUUID string) (interface{}, error)
	// GetPaymentList()
}

func New(username, password, endpoint, callback, invoiceCode, merchantId string) QPay {
	return &qpay{
		endpoint:    endpoint,
		password:    password,
		username:    username,
		callback:    callback,
		invoiceCode: invoiceCode,
		merchantId:  merchantId,
		loginObject: nil,
	}
}

func (q *qpay) CreateInvoice(input QPayCreateInvoiceInput) (QPaySimpleInvoiceResponse, error) {
	vals := url.Values{}
	for k, v := range input.CallbackParam {
		vals.Add(k, v)
	}

	amountInt := int64(input.Amount)
	request := QPaySimpleInvoiceRequest{
		InvoiceCode:         q.invoiceCode,
		SenderInvoiceCode:   input.SenderCode,
		InvoiceReceiverCode: input.ReceiverCode,
		InvoiceDescription:  input.Description,
		Amount:              amountInt,
		CallbackUrl:         fmt.Sprintf("%s?%s", q.callback, vals.Encode()),
	}

	res, err := q.httpRequestQPay(request, QPayInvoiceCreate, "")
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
func (q *qpay) CancelInvoice(invoiceId string) (interface{}, error) {
	res, err := q.httpRequestQPay(nil, QPayInvoiceCancel, invoiceId)
	if err != nil {
		return nil, err
	}

	var response interface{}
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpay) GetPayment(invoiceId string) (interface{}, error) {
	res, err := q.httpRequestQPay(nil, QPayPaymentGet, invoiceId)
	if err != nil {
		return nil, err
	}

	var response interface{}
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpay) CheckPayment(invoiceId string, pageLimit, pageNumber int32) (QpayPaymentCheckResponse, error) {
	req := QpayPaymentCheckRequest{}
	req.ObjectID = invoiceId
	req.ObjectType = "INVOICE"
	req.Offset.PageLimit = pageLimit
	req.Offset.PageNumber = pageNumber

	var response QpayPaymentCheckResponse

	res, err := q.httpRequestQPay(req, QPayPaymentCheck, "")
	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpay) CancelPayment(invoiceId, paymentUUID string) (QpayPaymentCheckResponse, error) {
	var req QpayPaymentCancelRequest

	req.CallbackUrl = q.callback + paymentUUID
	req.Note = "Cancel payment - " + invoiceId

	var response QpayPaymentCheckResponse

	res, err := q.httpRequestQPay(req, QPayPaymentCancel, invoiceId)
	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}

func (q *qpay) RefundPayment(invoiceId, paymentUUID string) (interface{}, error) {
	var req QpayPaymentCancelRequest

	req.CallbackUrl = q.callback + paymentUUID
	req.Note = "Cancel payment - " + invoiceId

	var response interface{}

	res, err := q.httpRequestQPay(req, QPayPaymentRefund, invoiceId)
	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}

// func (q *qpay) GetPaymentList() (QpayPaymentListRequest, error) {
// 	var req QpayPaymentListRequest
// 	req.MerchantID = q.merchantId

// 	res, err := utils.HttpRequestQpay(list, helper.QPayPaymentList, "")
// 	if err != nil {
// 		return res, err
// 	}
// }
