package ebarimt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ebarimt struct {
	customerNo string
	endpoint   string
}

type Ebarimt interface {
	GetNewEBarimt(*CreateEbarimtInput) (*CreateEbarimtResponse, error)
	CheckApi() error
	ReturnBill(billId, date string) (bool, error)
	SendData() error
}

func New(endpoint string) Ebarimt {
	return ebarimt{
		endpoint: endpoint,
	}
}

func float64ToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func stockInputToStock(input []StockInput) (stocks []Stock, amount float64, vat float64, citytax float64) {
	amount = 0
	vat = 0
	citytax = 0

	for _, v := range input {
		amount += v.UnitPrice * v.Qty
		vat += v.Vat
		citytax += v.CityTax
		stocks = append(stocks, Stock{
			Code:        v.Code,
			Name:        v.Name,
			Qty:         float64ToString(v.Qty),
			MeasureUnit: v.MeasureUnit,
			UnitPrice:   float64ToString(v.UnitPrice),
			CityTax:     float64ToString(v.CityTax),
			Vat:         float64ToString(v.Vat),
			BarCode:     v.BarCode,
			TotalAmount: float64ToString(v.UnitPrice * v.Qty),
		})
	}
	return
}

func createInputToRequestBody(input CreateEbarimtInput) *CreateEbarimtRequest {
	if input.DistrictCode == "" {
		input.DistrictCode = "34"
	}
	if input.BranchNo == "" {
		input.BranchNo = "001"
	}
	stocks, amount, vat, citytax := stockInputToStock(input.Stocks)
	return &CreateEbarimtRequest{
		Amount:        float64ToString(amount),
		Vat:           float64ToString(vat),
		CashAmount:    float64ToString(0),
		NonCashAmount: float64ToString(amount),
		CityTax:       float64ToString(citytax),
		CustomerNo:    input.CustomerNo,
		BillType:      string(input.BillType),
		BranchNo:      input.BranchNo,
		DistrictCode:  input.DistrictCode,
		Stocks:        stocks,
	}
}

func (b ebarimt) GetNewEBarimt(bodyraw *CreateEbarimtInput) (*CreateEbarimtResponse, error) {
	body := createInputToRequestBody(*bodyraw)
	if bodyraw.BillType == EBarimtOrganizationType && body.CustomerNo == "" {
		return nil, errors.New("CustomerNo is required")
	}

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		body.CustomerNo = b.customerNo
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}
	req, err := http.NewRequest("POST", b.endpoint+"/put", requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var responseBody CreateEbarimtResponse
	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func (b ebarimt) SendData() error {
	resp, err := http.Get(b.endpoint + "sendData")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (b ebarimt) ReturnBill(billId, date string) (bool, error) {
	var resp *http.Response
	var err error
	var url string
	var body []byte

	url = b.endpoint + "/returnBill"
	body, err = json.Marshal(map[string]string{
		"billId": billId,
		"date":   date,
	})
	if err != nil {
		return false, err
	}
	resp, err = http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 200 {
		return false, errors.New("return bill failed")
	}

	var responseBody ReturnBillResponse
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return false, err
	}

	return responseBody.Success, nil
}

func (b ebarimt) CheckApi() error {
	_, err := http.Get(b.endpoint)
	return err
}
