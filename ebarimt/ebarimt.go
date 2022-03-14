package ebarimt

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ebarimt struct {
	customerNo string
	endpoint   string
}

func New(customerNo, endpoint string) Ebarimt {
	return ebarimt{
		customerNo: customerNo,
		endpoint:   endpoint,
	}
}

type Ebarimt interface {
	GetNewEBarimt(*CreateEbarimtRequest) (*CreateEbarimtResponse, error)
	CheckApi() error
	ReturnBill(billId, date string) (bool, error)
	SendData() error
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

func (b ebarimt) GetNewEBarimt(body *CreateEbarimtRequest) (*CreateEbarimtResponse, error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}
	req, err := http.NewRequest("POST", b.endpoint+"/put", requestBody)
	if err != nil {
		err = errors.New("НӨАТ хүсэтийг боловсруулж чадсангүй")
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.New("НӨАТ хүсэтийг боловсруулж чадсангүй")
		return nil, err
	}

	var responseBody CreateEbarimtResponse
	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		err = errors.New("НӨАТ хүсэтийг боловсруулж чадсангүй")
		return nil, err
	}

	return &responseBody, nil
}
