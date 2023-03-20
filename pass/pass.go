package pass

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type pass struct {
	endpoint       string
	ecommerceToken string
	callback       string
}

type Pass interface {
	CreateOrder(amount int64, callbackParams map[string]string) (CreateOrderResponse, error)
	InqueryOrder(orderID string) (OrderInqueryResponse, error)
	NotifyOrder(orderID, phone string) (OrderNotifyResponse, error)
	CancelOrder(orderID string) (OrderCancelResponse, error)
	VoidOrder(orderID string) (OrderVoidResponse, error)
}

func New(endpoint, ecommerceToken, callback string) Pass {
	return &pass{
		endpoint:       endpoint,
		ecommerceToken: ecommerceToken,
		callback:       callback,
	}
}

func (p *pass) CreateOrder(amount int64, callbackParams map[string]string) (CreateOrderResponse, error) {
	vals := url.Values{}
	for k, v := range callbackParams {
		vals.Add(k, v)
	}

	amountInt := amount * 100
	request := CreateOrderInput{
		Amount:         amountInt,
		EcommerceToken: p.ecommerceToken,
		CallbackURL:    fmt.Sprintf("%s?%s", p.callback, vals.Encode()),
	}

	res, err := p.httpRequestPass(request, PassCreateOrder, "")
	if err != nil {
		return CreateOrderResponse{}, err
	}

	var response CreateOrderResponse
	json.Unmarshal(res, &response)

	if response.Msg.toError() != nil {
		return CreateOrderResponse{}, response.Msg.toError()
	}
	return response, nil
}

func (p *pass) InqueryOrder(orderID string) (OrderInqueryResponse, error) {
	request := OrderInqueryInput{
		EcommerceToken: p.ecommerceToken,
		OrderID:        orderID,
	}

	res, err := p.httpRequestPass(request, PassInqueryOrder, "")
	if err != nil {
		return OrderInqueryResponse{}, err
	}

	var response OrderInqueryResponse
	json.Unmarshal(res, &response)

	if response.Msg.toError() != nil {
		return OrderInqueryResponse{}, response.Msg.toError()
	}
	return response, nil
}

func (p *pass) NotifyOrder(orderID, phone string) (OrderNotifyResponse, error) {
	request := OrderNotifyInput{
		EcommerceToken: p.ecommerceToken,
		OrderID:        orderID,
		Phone:          phone,
	}

	res, err := p.httpRequestPass(request, PassNotifyOrder, "")
	if err != nil {
		return OrderNotifyResponse{}, err
	}

	var response OrderNotifyResponse
	json.Unmarshal(res, &response)

	if response.Msg.toError() != nil {
		return OrderNotifyResponse{}, response.Msg.toError()
	}
	return response, nil
}

func (p *pass) CancelOrder(orderID string) (OrderCancelResponse, error) {
	request := OrderCancelInput{
		EcommerceToken: p.ecommerceToken,
		OrderID:        orderID,
	}

	res, err := p.httpRequestPass(request, PassCancelOrder, "")
	if err != nil {
		return OrderCancelResponse{}, err
	}

	var response OrderCancelResponse
	json.Unmarshal(res, &response)

	if response.Msg.toError() != nil {
		return OrderCancelResponse{}, response.Msg.toError()
	}
	return response, nil
}

func (p *pass) VoidOrder(orderID string) (OrderVoidResponse, error) {
	request := OrderVoidInput{
		EcommerceToken: p.ecommerceToken,
		OrderID:        orderID,
	}

	res, err := p.httpRequestPass(request, PassVoidOrder, "")
	if err != nil {
		return OrderVoidResponse{}, err
	}

	var response OrderVoidResponse
	json.Unmarshal(res, &response)

	if response.Msg.toError() != nil {
		return OrderVoidResponse{}, response.Msg.toError()
	}
	return response, nil
}
