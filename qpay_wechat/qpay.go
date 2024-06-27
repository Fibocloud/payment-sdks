package qpay_wechat

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type qpay_auth struct {
	endpoint    string
	password    string
	username    string
	callback    string
	invoiceCode string
	merchantId  string
	loginObject *qpayLoginResponse
}

type QPayAuth interface {
	CreateInvoice(input QPayCreateInvoiceInput) (QPaySimpleInvoiceResponse, QPayAuth, error)
	GetInvoice(invoiceId string) (QpayInvoiceGetResponse, QPayAuth, error)
	CancelInvoice(invoiceId string) (interface{}, QPayAuth, error)
	GetPayment(invoiceId string) (interface{}, QPayAuth, error)
	CheckPayment(invoiceId string, pageLimit, pageNumber int64) (QpayPaymentCheckResponse, QPayAuth, error)
	CancelPayment(invoiceId, paymentUUID string) (QpayPaymentCheckResponse, QPayAuth, error)
	RefundPayment(invoiceId, paymentUUID string) (interface{}, QPayAuth, error)
	// GetPaymentList()
}

func New(username, password, endpoint, callback, invoiceCode, merchantId string) QPayAuth {
	return &qpay_auth{
		endpoint:    endpoint,
		password:    password,
		username:    username,
		callback:    callback,
		invoiceCode: invoiceCode,
		merchantId:  merchantId,
		loginObject: func() *qpayLoginResponse {
			authObj, authErr := authQPayV2(username, password, endpoint, callback, invoiceCode, merchantId)
			if authErr != nil {
				// err = authErr
				return &qpayLoginResponse{}
			}
			return &authObj
		}(),
	}
}

func (q *qpay_auth) CreateInvoice(input QPayCreateInvoiceInput) (QPaySimpleInvoiceResponse, QPayAuth, error) {
	vals := url.Values{}
	for k, v := range input.CallbackParam {
		vals.Add(k, v)
	}

	amountInt := int64(input.Amount)
	request := QPaySimpleInvoiceRequest{
		InvoiceCode:         q.invoiceCode,
		SenderInvoiceCode:   input.SenderCode,
		SenderBranchCode:    input.SenderBranchCode,
		InvoiceReceiverCode: input.ReceiverCode,
		InvoiceDescription:  input.Description,
		Amount:              amountInt,
		CallbackUrl:         fmt.Sprintf("%s?%s", q.callback, vals.Encode()),
	}

	res, err := q.httpRequest(request, QPayInvoiceCreate, "")
	if err != nil {
		return QPaySimpleInvoiceResponse{}, q, err
	}

	var response QPaySimpleInvoiceResponse
	json.Unmarshal(res, &response)

	return response, q, nil
}
func (q *qpay_auth) GetInvoice(invoiceId string) (QpayInvoiceGetResponse, QPayAuth, error) {
	res, err := q.httpRequest(nil, QPayInvoiceGet, invoiceId)
	if err != nil {
		return QpayInvoiceGetResponse{}, q, err
	}

	var response QpayInvoiceGetResponse
	json.Unmarshal(res, &response)

	return response, q, nil
}
func (q *qpay_auth) CancelInvoice(invoiceId string) (interface{}, QPayAuth, error) {
	res, err := q.httpRequest(nil, QPayInvoiceCancel, invoiceId)
	if err != nil {
		return nil, q, err
	}

	var response interface{}
	json.Unmarshal(res, &response)

	return response, q, nil
}

func (q *qpay_auth) GetPayment(invoiceId string) (interface{}, QPayAuth, error) {
	res, err := q.httpRequest(nil, QPayPaymentGet, invoiceId)
	if err != nil {
		return nil, q, err
	}

	var response interface{}
	json.Unmarshal(res, &response)

	return response, q, nil
}

func (q *qpay_auth) CheckPayment(invoiceId string, pageLimit, pageNumber int64) (QpayPaymentCheckResponse, QPayAuth, error) {
	req := QpayPaymentCheckRequest{}
	req.ObjectID = invoiceId
	req.ObjectType = "INVOICE"
	req.Offset.PageLimit = pageLimit
	req.Offset.PageNumber = pageNumber

	var response QpayPaymentCheckResponse

	res, err := q.httpRequest(req, QPayPaymentCheck, "")
	if err != nil {
		return response, q, err
	}

	json.Unmarshal(res, &response)

	return response, q, nil
}

func (q *qpay_auth) CancelPayment(invoiceId, paymentUUID string) (QpayPaymentCheckResponse, QPayAuth, error) {
	var req QpayPaymentCancelRequest

	req.CallbackUrl = q.callback + paymentUUID
	req.Note = "Cancel payment - " + invoiceId

	var response QpayPaymentCheckResponse

	res, err := q.httpRequest(req, QPayPaymentCancel, invoiceId)
	// ret := func() QPayAuth {
	// 	return &qpay_auth{
	// 		endpoint:    q.endpoint,
	// 		password:    q.password,
	// 		username:    q.username,
	// 		callback:    q.callback,
	// 		invoiceCode: q.invoiceCode,
	// 		merchantId:  q.merchantId,
	// 		loginObject: q.loginObject,
	// 	}
	// }()
	if err != nil {
		return response, q, err
	}

	json.Unmarshal(res, &response)

	return response, q, nil
}

func (q *qpay_auth) RefundPayment(invoiceId, paymentUUID string) (interface{}, QPayAuth, error) {
	var req QpayPaymentCancelRequest

	req.CallbackUrl = q.callback + paymentUUID
	req.Note = "Cancel payment - " + invoiceId

	var response interface{}

	res, err := q.httpRequest(req, QPayPaymentRefund, invoiceId)
	if err != nil {
		return response, q, err
	}

	json.Unmarshal(res, &response)

	return response, q, nil
}

// func (q *qpay) GetPaymentList() (QpayPaymentListRequest, error) {
// 	var req QpayPaymentListRequest
// 	req.MerchantID = q.merchantId

// 	res, err := utils.HttpRequestQpay(list, helper.QPayPaymentList, "")
// 	if err != nil {
// 		return res, err
// 	}
// }
