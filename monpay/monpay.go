package monpay

import (
	"encoding/json"
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
	// GenerateQR(input MchatOnlineQrGenerateRequest) (MchatOnlineQrGenerateResponse, error)
	// CheckQR(qr string) (MchatOnlineQrCheckResponse, error)
}

func New(endpoint, username, accountId string) Monpay {
	return monpay{
		endpoint:  endpoint,
		username:  username,
		accoutnId: accountId,
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
	return
}

func (m monpay) CheckQr(uuid string) (response MonpayCheckResponse, err error) {
	res, err := m.httpRequestMonpay(nil, MonpayCheckQr, uuid)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

var decoder = schema.NewDecoder()

func CallbackParser(url *url.URL) (response MonpayCallback) {
	err := decoder.Decode(&response, url.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	} else {
		log.Println("GET parameters : ", response)
	}
	return
}
