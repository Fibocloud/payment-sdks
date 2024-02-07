package tokipaythirdparty

import (
	"encoding/json"
	"errors"
)

type tokipayThirdParty struct {
	endpoint      string
	authorization string
	merchantId    string
	SuccessUrl    string
	FailureUrl    string
}

type TokipayThirdParty interface {
	PaymentSentUser(input TokipayPaymentInput) (TokipayPaymentResponse, error)
	PaymentStatus(requestId string) (TokipayPaymentStatusResponse, error)
}

func New(endpoint, authorization, merchantId, successUrl, failureUrl string) TokipayThirdParty {
	return &tokipayThirdParty{
		endpoint:      endpoint,
		authorization: authorization,
		merchantId:    merchantId,
		SuccessUrl:    successUrl,
		FailureUrl:    failureUrl,
	}
}

func (q *tokipayThirdParty) PaymentSentUser(input TokipayPaymentInput) (TokipayPaymentResponse, error) {
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

func (q *tokipayThirdParty) PaymentStatus(requestId string) (TokipayPaymentStatusResponse, error) {

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
