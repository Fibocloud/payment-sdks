package hipay

import (
	"encoding/json"
	"errors"
	"fmt"
)

type hipay struct {
	endpoint string
	token    string
	entityId string
}

type Hipay interface {
	Checkout(amount float64) (HipayCheckoutResponse, error)
	CheckoutGet(checkoutId string) (HipayCheckoutGetResponse, error)
	PaymentGet(paymentId string) (HipayPaymentGetResponse, error)
	PaymentCorrection(paymentId string) (HipayPaymentCorrectionResponse, error)
	Statement(date string) (HipayStatementResponse, error)
}

func New(endpoint, token, entityId string) Hipay {
	return &hipay{
		endpoint: endpoint,
		token:    token,
		entityId: entityId,
	}
}

func (h *hipay) Checkout(amount float64) (HipayCheckoutResponse, error) {
	request := HipayCheckoutRequest{
		EntityID: h.entityId,
		Amount:   amount,
		Currency: "MNT",
		QrData:   true,
		Signal:   false,
	}
	res, err := h.httpRequestHipay(request, HipayCheckout, "")
	if err != nil {
		return HipayCheckoutResponse{}, err
	}

	var response HipayCheckoutResponse
	if err := json.Unmarshal(res, &response); err != nil {
		fmt.Println(err.Error())
		return HipayCheckoutResponse{}, err
	}
	if response.Code != 1 {
		return HipayCheckoutResponse{}, errors.New(response.Description + ": " + response.Details[0].Field + " - " + response.Details[0].Issue)
	}
	return response, nil
}

func (h *hipay) CheckoutGet(checkoutId string) (HipayCheckoutGetResponse, error) {
	ext := checkoutId + "?entityId=" + h.entityId
	res, err := h.httpRequestHipay(nil, HipayCheckoutGet, ext)
	if err != nil {
		return HipayCheckoutGetResponse{}, err
	}
	var response HipayCheckoutGetResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return HipayCheckoutGetResponse{}, err
	}
	if response.Code != 1 {
		return HipayCheckoutGetResponse{}, errors.New(response.Description + ": " + response.Details[0].Field + " - " + response.Details[0].Issue)
	}
	return response, nil
}

func (h *hipay) PaymentGet(paymentId string) (HipayPaymentGetResponse, error) {
	ext := paymentId + "?entityId=" + h.entityId
	res, err := h.httpRequestHipay(nil, HipayPaymentGet, ext)
	if err != nil {
		return HipayPaymentGetResponse{}, err
	}
	var response HipayPaymentGetResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return HipayPaymentGetResponse{}, err
	}
	if response.Code != 1 {
		return HipayPaymentGetResponse{}, errors.New(response.Description + ": " + response.Details[0].Field + " - " + response.Details[0].Issue)
	}
	return response, nil
}

func (h *hipay) PaymentCorrection(paymentId string) (HipayPaymentCorrectionResponse, error) {
	body := HipayPaymentCorrectionRequest{
		EntityID:  h.entityId,
		PaymentID: paymentId,
	}
	res, err := h.httpRequestHipay(body, HipayPaymentCorrection, "")
	if err != nil {
		return HipayPaymentCorrectionResponse{}, err
	}
	var response HipayPaymentCorrectionResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return HipayPaymentCorrectionResponse{}, err
	}
	if response.Code != 1 {
		return HipayPaymentCorrectionResponse{}, errors.New(response.Description + ": " + response.Details[0].Field + " - " + response.Details[0].Issue)
	}
	return response, nil
}

func (h *hipay) Statement(date string) (HipayStatementResponse, error) {
	body := HipayStatementRequest{
		EntityID: h.entityId,
		Date:     date,
	}
	res, err := h.httpRequestHipay(body, HipayStatement, "")
	if err != nil {
		return HipayStatementResponse{}, err
	}
	var response HipayStatementResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return HipayStatementResponse{}, err
	}
	if response.Code != 1 {
		return HipayStatementResponse{}, errors.New(response.Description + ": " + response.Details[0].Field + " - " + response.Details[0].Issue)
	}
	return response, nil
}
