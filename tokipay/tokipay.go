package tokipay

import (
	"encoding/json"
	"errors"
)

type tokipay struct {
	endpoint      string
	imApiKey      string
	authorization string
	merchantId    string
	SuccessUrl    string
	FailureUrl    string
	AppSchemaIos  string
}

type Tokipay interface {
	PaymentQr(input TokipayPaymentInput) (TokipayPaymentResponse, error)
	PaymentSentUser(input TokipayPaymentInput) (TokipayPaymentResponse, error)
	PaymentScanUser(input TokipayPaymentInput) (TokipayPaymentResponse, error)
	PaymentStatus(requestId string) (TokipayPaymentStatusResponse, error)
	PaymentCancel(requestId string) (TokipayPaymentStatusResponse, error)
	PaymentThirdPartyDeeplink(input TokipayPaymentInput) (TokipayDeeplinkResponse, error)
	PaymentThirdPartyPhoneRequest(input TokipayPaymentInput) (TokipayThirdPartyPhoneResponse, error)
	PaymentThirdPartyStatus(requestId string) (TokipayPaymentStatusResponse, error)
}

func New(endpoint, imApiKey, authorization, merchantId, successUrl, failureUrl, appSchemaIos string) Tokipay {
	return &tokipay{
		endpoint:      endpoint,
		imApiKey:      imApiKey,
		authorization: authorization,
		merchantId:    merchantId,
		SuccessUrl:    successUrl,
		FailureUrl:    failureUrl,
		AppSchemaIos:  appSchemaIos,
	}
}

func (q *tokipay) PaymentQr(input TokipayPaymentInput) (TokipayPaymentResponse, error) {
	request := TokipayPaymentQrRequest{
		SuccessUrl:    q.SuccessUrl,
		FailureUrl:    q.FailureUrl,
		OrderId:       input.OrderId,
		MerchantId:    q.merchantId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		Authorization: q.authorization,
	}

	res, err := q.httpRequestTokipayPOS(request, TokipayPaymentQr, "")
	if err != nil {
		return TokipayPaymentResponse{}, err
	}
	var response TokipayPaymentResponse
	json.Unmarshal(res, &response)
	if response.StatusCode != 200 {
		return TokipayPaymentResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentSentUser(input TokipayPaymentInput) (TokipayPaymentResponse, error) {
	request := TokipayPaymentSentUserRequest{
		SuccessUrl:    q.SuccessUrl,
		FailureUrl:    q.FailureUrl,
		OrderId:       input.OrderId,
		MerchantId:    q.merchantId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		Authorization: q.authorization,
		PhoneNo:       input.PhoneNo,
		CountryCode:   input.CountryCode,
	}

	res, err := q.httpRequestTokipayPOS(request, TokipayPaymentSentUser, "")
	if err != nil {
		return TokipayPaymentResponse{}, err
	}

	var response TokipayPaymentResponse
	json.Unmarshal(res, &response)
	if response.StatusCode != 200 {
		return TokipayPaymentResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentScanUser(input TokipayPaymentInput) (TokipayPaymentResponse, error) {
	request := TokipayPaymentScanUserRequest{
		SuccessUrl:    q.SuccessUrl,
		FailureUrl:    q.FailureUrl,
		OrderId:       input.OrderId,
		MerchantId:    q.merchantId,
		Amount:        input.Amount,
		Notes:         input.Notes,
		Authorization: q.authorization,
		RequestId:     input.RequestId,
	}

	res, err := q.httpRequestTokipayPOS(request, TokipayPaymentScanUser, "")
	if err != nil {
		return TokipayPaymentResponse{}, err
	}

	var response TokipayPaymentResponse
	json.Unmarshal(res, &response)

	if response.StatusCode != 200 {
		return TokipayPaymentResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentStatus(requestId string) (TokipayPaymentStatusResponse, error) {

	res, err := q.httpRequestTokipayPOS(nil, TokipayPaymentStatus, requestId)
	if err != nil {
		return TokipayPaymentStatusResponse{}, err
	}

	var response TokipayPaymentStatusResponse
	json.Unmarshal(res, &response)
	if response.StatusCode != 200 {
		return TokipayPaymentStatusResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentThirdPartyDeeplink(input TokipayPaymentInput) (TokipayDeeplinkResponse, error) {
	request := TokipayDeeplinkRequest{
		SuccessUrl: func() string {
			if input.SuccessUrl != "" {
				return input.SuccessUrl
			}
			return q.SuccessUrl
		}(),
		FailureUrl:        q.FailureUrl,
		OrderId:           input.OrderId,
		MerchantId:        q.merchantId,
		Amount:            input.Amount,
		Notes:             input.Notes,
		AppSchemaIos:      q.AppSchemaIos,
		Authorization:     q.authorization,
		TokiWebSuccessUrl: q.SuccessUrl,
		TokiWebFailureUrl: q.FailureUrl,
	}

	res, err := q.httpRequestTokipayThirdParty(request, TokipayDeeplink, "")
	if err != nil {
		return TokipayDeeplinkResponse{}, err
	}

	var response TokipayDeeplinkResponse
	json.Unmarshal(res, &response)

	if response.StatusCode != 200 {
		return TokipayDeeplinkResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentThirdPartyPhoneRequest(input TokipayPaymentInput) (TokipayThirdPartyPhoneResponse, error) {
	request := TokipayThirdPartyPhoneRequest{
		SuccessUrl:        q.SuccessUrl,
		FailureUrl:        q.FailureUrl,
		OrderId:           input.OrderId,
		MerchantId:        q.merchantId,
		Amount:            input.Amount,
		Notes:             input.Notes,
		PhoneNo:           input.PhoneNo,
		CountryCode:       "+976",
		Authorization:     q.authorization,
		TokiWebSuccessUrl: q.SuccessUrl,
		TokiWebFailureUrl: q.FailureUrl,
	}

	res, err := q.httpRequestTokipayThirdParty(request, TokipayPhoneRequest, "")
	if err != nil {
		return TokipayThirdPartyPhoneResponse{}, err
	}

	var response TokipayThirdPartyPhoneResponse
	json.Unmarshal(res, &response)

	if response.StatusCode != 200 {
		return TokipayThirdPartyPhoneResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentThirdPartyStatus(requestId string) (TokipayPaymentStatusResponse, error) {

	res, err := q.httpRequestTokipayThirdParty(nil, TokipayTransactionStatus, requestId)
	if err != nil {
		return TokipayPaymentStatusResponse{}, err
	}

	var response TokipayPaymentStatusResponse
	json.Unmarshal(res, &response)
	if response.StatusCode != 200 {
		return TokipayPaymentStatusResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentCancel(requestId string) (TokipayPaymentStatusResponse, error) {

	res, err := q.httpRequestTokipayPOS(nil, TokipayPaymentCancel, requestId)
	if err != nil {
		return TokipayPaymentStatusResponse{}, err
	}

	var response TokipayPaymentStatusResponse
	json.Unmarshal(res, &response)
	if response.StatusCode != 200 {
		return TokipayPaymentStatusResponse{}, errors.New(response.Error + ":" + response.Message)
	}
	return response, nil
}

func (q *tokipay) PaymentRefund(input TokipayRefundInput) (TokipayPaymentResponseExt, error) {

	request := TokipayRefundRequest{
		RequestId:    input.RequestId,
		RefundAmount: input.RefundAmount,
		MerchantId:   q.merchantId,
	}
	res, err := q.httpRequestTokipayPOS(request, TokipayRefund, "")
	if err != nil {
		return TokipayPaymentResponseExt{}, err
	}

	var response TokipayPaymentResponseExt
	json.Unmarshal(res, &response)

	return response, nil
}
