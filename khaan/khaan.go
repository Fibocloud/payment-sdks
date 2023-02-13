package khaan

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

type (
	Khaan interface {
		RegisterOrder(OrderRegisterInput) (*RegisterOrderResponse, error)
		CheckOrder(string) (*OrderStatusResponse, error)
	}

	khaan struct {
		endPoint string
		username string
		password string
		langauge string // language codes like en, mn etc... https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
	}
)

func New(endpoint, username, password, language string) Khaan {
	return khaan{
		endPoint: endpoint,
		username: username,
		password: password,
		langauge: language,
	}
}

type OrderRegisterInput struct {
	OrderNumber     string
	Amount          float64
	SuccessCallback string
	FailCallback    string
}

type KhaanRequestBody struct {
	Username string `json:"userName"`
	Password string `json:"Password"`
	Language string `json:"language"`
}

type OrderRequest struct {
	OrderNumber     string            `json:"orderNumber"`
	Amount          string            `json:"amount"`
	SuccessCallback string            `json:"returnUrl"`
	FailCallback    string            `json:"failUrl"`
	JsonParams      map[string]string `json:"jsonParams"`
	KhaanRequestBody
}

type RegisterOrderResponse struct {
	OrderId string
	FormUrl string
}

type orderStatus struct {
	OrderStatus  string `json:"orderStatus"`
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
	OrderNumber  string `json:"OrderNumber"`
	Ip           string `json:"Ip"`
}

type OrderStatusResponse struct {
	Success      bool
	ErrorCode    string
	ErrorMessage string
	OrderNumber  string
	Ip           string
}

type OrderStatusRequest struct {
	KhaanRequestBody
	OrderId string `json:"orderId"`
}

func (kh khaan) CheckOrder(orderId string) (*OrderStatusResponse, error) {
	req := OrderStatusRequest{
		KhaanRequestBody: KhaanRequestBody{
			Username: kh.username,
			Password: kh.password,
			Language: kh.langauge,
		},
	}
	resp, err := sendRequestKhaanEndpoint(kh.endPoint, KhaanOrderRegister, req)
	if err != nil {
		return nil, err
	}

	var checkRes orderStatus
	err = json.Unmarshal(resp, &checkRes)
	if err != nil {
		return nil, err
	}

	var res OrderStatusResponse
	res.ErrorCode = checkRes.ErrorCode
	res.ErrorMessage = checkRes.ErrorMessage
	if checkRes.OrderStatus == "2" {
		res.Success = true
	} else {
		res.Success = false
	}
	return &res, err
}

func (kh khaan) RegisterOrder(input OrderRegisterInput) (*RegisterOrderResponse, error) {
	req := OrderRequest{
		OrderNumber:     input.OrderNumber,
		Amount:          fmt.Sprintf("%.2f", input.Amount),
		SuccessCallback: input.SuccessCallback,
		FailCallback:    input.FailCallback,
		JsonParams: map[string]string{
			"orderNumber": input.OrderNumber,
		},
		KhaanRequestBody: KhaanRequestBody{
			Username: kh.username,
			Password: kh.password,
			Language: kh.langauge,
		},
	}

	res, err := sendRequestKhaanEndpoint(kh.endPoint, KhaanOrderRegister, req)
	if err != nil {
		return nil, err
	}

	var orderRes RegisterOrderResponse
	err = json.Unmarshal(res, &orderRes)
	if err != nil {
		return nil, err
	}

	return &orderRes, err
}

func sendRequestKhaanEndpoint(endpoint string, api utils.API, body interface{}) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, endpoint+api.Url, requestBody)
	req.Header.Add("Content-Type", utils.HttpContent)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	response, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unknown error occurred")
	}
	return
}
