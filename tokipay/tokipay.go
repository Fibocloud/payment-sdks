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
	PaymentStatus(requestId string) (TokipayPaymentStatusResponse, error)
	PaymentCancel(requestId string) (TokipayPaymentStatusResponse, error)
}

func New(endpoint, apiKey, imApiKey, authorization, merchantId string) Tokipay {
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
		SuccessUrl:    "www.mtm.mn",
		FailureUrl:    "www.mtm.mn",
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
		SuccessUrl:    "www.mtm.mn",
		FailureUrl:    "www.mtm.mn",
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
		SuccessUrl:    "www.mtm.mn",
		FailureUrl:    "www.mtm.mn",
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

func (q *tokipay) PaymentStatus(requestId string) (TokipayPaymentStatusResponse, error) {

	res, err := q.httpRequestTokipay(nil, TokipayPaymentStatus, requestId)
	if err != nil {
		return TokipayPaymentStatusResponse{}, err
	}

	var response TokipayPaymentStatusResponse
	json.Unmarshal(res, &response)

	return response, nil
}

func (q *tokipay) PaymentCancel(requestId string) (TokipayPaymentStatusResponse, error) {

	res, err := q.httpRequestTokipay(nil, TokipayPaymentCancel, requestId)
	if err != nil {
		return TokipayPaymentStatusResponse{}, err
	}

	var response TokipayPaymentStatusResponse
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
