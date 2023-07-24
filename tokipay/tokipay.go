package tokipay

import (
	"encoding/json"
)

type tokipay struct {
	endpoint      string
	apiKey        string
	imApiKey      string
	authorization string
	merchantId    string
}

type Tokipay interface {
	PaymentQr(input TokipayPaymentQrInput) (TokipayPaymentResponse, error)
	PaymentSentUser(input TokipayPaymentQrInput) (TokipayPaymentResponse, error)
	PaymentScanUser(input TokipayPaymentQrInput) (TokipayPaymentResponse, error)
}

func New(endpoint, apiKey, imApiKey, authorization, merchantId, successUrl, failureUrl string) Tokipay {
	return &tokipay{
		endpoint:      endpoint,
		apiKey:        apiKey,
		imApiKey:      imApiKey,
		authorization: authorization,
		merchantId:    merchantId,
	}
}

func (q *tokipay) PaymentQr(input TokipayPaymentQrInput) (TokipayPaymentResponse, error) {
	request := TokipayPaymentQrRequest{
		OrderId:       input.OrderId,
		MerchantId:    q.merchantId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		Authorization: q.authorization,
	}

	res, err := q.httpRequestTokipay(request, TokipayPaymentQr, "")
	if err != nil {
		return TokipayPaymentResponse{}, err
	}

	var response TokipayPaymentResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *tokipay) PaymentSentUser(input TokipayPaymentQrInput) (TokipayPaymentResponse, error) {
	request := TokipayPaymentSentUserRequest{
		OrderId:       input.OrderId,
		MerchantId:    q.merchantId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		Authorization: q.authorization,
		PhoneNo:       input.PhoneNo,
		CountryCode:   input.CountryCode,
	}

	res, err := q.httpRequestTokipay(request, TokipayPaymentSentUser, "")
	if err != nil {
		return TokipayPaymentResponse{}, err
	}

	var response TokipayPaymentResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *tokipay) PaymentScanUser(input TokipayPaymentQrInput) (TokipayPaymentResponse, error) {
	request := TokipayPaymentScanUserRequest{
		OrderId:       input.OrderId,
		MerchantId:    q.merchantId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		Authorization: q.authorization,
		RequestId:     input.RequestId,
	}

	res, err := q.httpRequestTokipay(request, TokipayPaymentScanUser, "")
	if err != nil {
		return TokipayPaymentResponse{}, err
	}

	var response TokipayPaymentResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *tokipay) PaymentStatus(requestId string) (TokipayPaymentResponseExt, error) {

	res, err := q.httpRequestTokipay(nil, TokipayPaymentStatus, requestId)
	if err != nil {
		return TokipayPaymentResponseExt{}, err
	}

	var response TokipayPaymentResponseExt
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *tokipay) PaymentCancel(requestId string) (TokipayPaymentResponseExt, error) {

	res, err := q.httpRequestTokipay(nil, TokipayPaymentCancel, requestId)
	if err != nil {
		return TokipayPaymentResponseExt{}, err
	}

	var response TokipayPaymentResponseExt
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *tokipay) PaymentRefund(input TokipayRefundInput) (TokipayPaymentResponseExt, error) {

	request := TokipayRefundRequest{
		RequestId:    input.RequestId,
		RefundAmount: input.RefundAmount,
		MerchantId:   q.merchantId,
	}
	res, err := q.httpRequestTokipay(request, TokipayRefund, "")
	if err != nil {
		return TokipayPaymentResponseExt{}, err
	}

	var response TokipayPaymentResponseExt
	json.Unmarshal(res, &response)

	return response, nil
}
