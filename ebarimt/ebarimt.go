package ebarimt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
)

type ebarimt struct {
	endpoint string
}

type ebarimtcli struct {
}

type Ebarimt interface {
	GetNewEBarimt(*CreateEbarimtInput) (*CreateEbarimtResponse, error)
	CheckApi() (*CheckResponse, error)
	ReturnBill(billId, date string) (bool, error)
	SendData() error
}

func New(endpoint string) Ebarimt {
	return ebarimt{
		endpoint: endpoint,
	}
}

func NewCli() Ebarimt {
	return ebarimtcli{}
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
		body.CustomerNo = bodyraw.CustomerNo
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

func (b ebarimtcli) GetNewEBarimt(bodyraw *CreateEbarimtInput) (*CreateEbarimtResponse, error) {
	body := createInputToRequestBody(*bodyraw)
	if bodyraw.BillType == EBarimtOrganizationType && body.CustomerNo == "" {
		return nil, errors.New("CustomerNo is required")
	}

	var requestByte []byte
	body.CustomerNo = bodyraw.CustomerNo
	requestByte, _ = json.Marshal(body)

	out, err := exec.Command("ebarimt", "put", fmt.Sprintf("'%s'", fmt.Sprintf("'%s'", string(requestByte)))).Output()
	if err != nil {
		return nil, err
	}
	fmt.Println("out :", string(out))
	var responseBody CreateEbarimtResponse
	err = json.NewDecoder(bytes.NewReader(out)).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func (b ebarimtcli) SendData() error {
	out, err := exec.Command("ebarimt", "send_data").Output()
	if err != nil {
		return err
	}
	fmt.Println("out :", string(out))
	return nil
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
		"returnBillId": billId,
		"date":         date,
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

func (b ebarimtcli) ReturnBill(billId, date string) (bool, error) {
	body, err := json.Marshal(map[string]string{
		"returnBillId": billId,
		"date":         date,
	})
	if err != nil {
		return false, err
	}
	out, err := exec.Command("ebarimt", "return_bill", fmt.Sprintf("'%s'", string(body))).Output()
	if err != nil {
		return false, err
	}
	fmt.Println("out :", string(out))

	rdr := bytes.NewReader(out)
	var responseBody ReturnBillResponse

	err = json.NewDecoder(rdr).Decode(&responseBody)
	if err != nil {
		return false, err
	}

	return responseBody.Success, nil
}

func (b ebarimt) CheckApi() (*CheckResponse, error) {
	var responseBody CheckResponse
	resp, err := http.Get(b.endpoint)
	if err != nil {
		return &responseBody, err
	}
	if resp.StatusCode != 200 {
		return &responseBody, errors.New("return bill failed")
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return &responseBody, err
	}
	return &responseBody, err

}

func (b ebarimtcli) CheckApi() (*CheckResponse, error) {
	var responseBody CheckResponse
	out, err := exec.Command("ebarimt", "check_api").Output()
	if err != nil {
		return nil, err
	}
	fmt.Println("out :", string(out))

	rdr := bytes.NewReader(out)

	err = json.NewDecoder(rdr).Decode(&responseBody)
	if err != nil {
		return nil, err
	}
	return &responseBody, nil
}
