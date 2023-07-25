package monpay

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"

	"github.com/gorilla/schema"
)

type monpay struct {
	endpoint    string
	username    string
	accoutnId   string
	callbackurl string
}

type Monpay interface {
	GenerateQr(input MonpayQrInput) (MonpayQrResponse, error)
	CheckQr(uuid string) (MonpayCheckResponse, error)
	CallbackParser(url *url.URL) MonpayCallback
}

func New(endpoint, username, accountId, callback string) Monpay {
	return monpay{
		endpoint:    endpoint,
		username:    username,
		accoutnId:   accountId,
		callbackurl: callback,
	}
}

func (m monpay) GenerateQr(input MonpayQrInput) (response MonpayQrResponse, err error) {
	invoice := MonpayQrRequest{
		Amount:       input.Amount,
		GenerateUUID: true,
		CallbackUrl:  m.callbackurl,
	}
	res, err := m.httpRequestMonpay(invoice, MonpayGenerateQr, "")
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	if response.Code != 0 {
		switch response.Code {
		case 5:
			err = errors.New("хүсэлт буруу")
		case 1:
			err = errors.New("unauthorized")
		case 999:
			err = errors.New("дотоод алдаа")
		default:
			err = errors.New("unknown error")
		}
	}
	return
}

func (m monpay) CheckQr(uuid string) (response MonpayCheckResponse, err error) {
	res, err := m.httpRequestMonpay(nil, MonpayCheckQr, uuid)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	if response.Code != 0 {
		switch response.Code {
		case 5:
			err = errors.New("хүсэлт буруу")
		case 23:
			err = errors.New("QR not scanned")
		case 1:
			err = errors.New("unauthorized")
		case 999:
			err = errors.New("дотоод алдаа")
		default:
			err = errors.New("unknown error")
		}
	}
	return
}

var decoder = schema.NewDecoder()

func (m monpay) CallbackParser(url *url.URL) (response MonpayCallback) {
	err := decoder.Decode(&response, url.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	} else {
		log.Println("GET parameters : ", response)
	}
	return
}
