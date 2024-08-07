package monpay

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type deeplink struct {
	endpoint     string
	webhookUrl   string
	redirectUrl  string
	clientId     string
	clientSecret string
	grantType    string
	accessToken  *AccessToken
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type Deeplink interface {
	// GetAccessToken() error
	CreateDeeplink(amount float64, invoiceType InvoiceType, branchUsername, desc, invoiceId string) (response DeeplinkCreateResponse, err error)
	CheckInvoice(invoiceID int) (response DeeplinkCheckResponse, err error)
	CallbackParser(url *url.URL) (response DeeplinkCallback)
}

func NewDeeplink(url, id, secret, grantType, webhookUrl, redirectUrl string) Deeplink {
	return &deeplink{
		endpoint:     url,
		clientId:     id,
		clientSecret: secret,
		grantType:    grantType,
		webhookUrl:   webhookUrl,
		redirectUrl:  redirectUrl,
		accessToken:  nil,
	}
}

// func (d deeplink) GetAccessToken() error {
// 	err := d.getAccessToken()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (d *deeplink) CreateDeeplink(amount float64, invoiceType InvoiceType, branchUsername, desc, invoiceId string) (response DeeplinkCreateResponse, err error) {
	body := DeeplinkCreateRequest{
		RedirectUri:      d.redirectUrl + "/" + invoiceId,
		ClientServiceUrl: d.webhookUrl + "/" + invoiceId,
		Amount:           amount,
		Receiver:         branchUsername,
		InvoiceType:      invoiceType,
		Description:      desc,
	}
	res, err := d.httpRequestDeeplink(body, MonpayDeeplinkCreate, "")
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return
	}
	return
}

func (d *deeplink) CheckInvoice(invoiceID int) (response DeeplinkCheckResponse, err error) {
	res, err := d.httpRequestDeeplink(nil, MonpayDeeplinkCheck, fmt.Sprintf("%d", invoiceID))
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return
	}
	return
}

func (d *deeplink) CallbackParser(url *url.URL) (response DeeplinkCallback) {
	err := decoder.Decode(&response, url.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	} else {
		log.Println("GET parameters : ", response)
	}
	return
}
