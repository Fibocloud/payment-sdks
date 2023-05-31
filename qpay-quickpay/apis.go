package qpayquick

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	QPayAuthToken = utils.API{
		Url:    "/auth/token",
		Method: http.MethodPost,
	}
	QPayAuthRefresh = utils.API{
		Url:    "/auth/refresh",
		Method: http.MethodPost,
	}

	QPayCreateCompany = utils.API{
		Url:    "/merchant/company",
		Method: http.MethodPost,
	}
	QPayCreatePerson = utils.API{
		Url:    "/merchant/person",
		Method: http.MethodPost,
	}
	QPayGetMerchant = utils.API{
		Url:    "/merchant/",
		Method: http.MethodGet,
	}
	QPayMerchantList = utils.API{
		Url:    "/merchant/list",
		Method: http.MethodPost,
	}

	QPayInvoiceCreate = utils.API{
		Url:    "/invoice",
		Method: http.MethodPost,
	}
	QPayInvoiceGet = utils.API{
		Url:    "/invoice/",
		Method: http.MethodGet,
	}

	QPayPaymentCheck = utils.API{
		Url:    "/payment/check",
		Method: http.MethodPost,
	}
)

// func (q *qpay) ExpireTokenForce() {
// 	q.loginObject.ExpiresIn = 0
// }

func (q *qpayquick) httpRequestQPay(body interface{}, api utils.API, urlExt string) (response []byte, err error) {
	// if q.loginObject == nil || time.Now().Unix() > int64(q.loginObject.RefreshExpiresIn) {

	authObj, authErr := q.authQPayV2()
	if authErr != nil {
		err = authErr
		return
	}

	q.loginObject = &authObj

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, q.endpoint+api.Url+urlExt, requestBody)
	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+q.loginObject.AccessToken)

	res, err := http.DefaultClient.Do(req)

	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}

	return
}

// AuthQPayV2 [Login to qpay]
func (q *qpayquick) authQPayV2() (authRes qpayLoginResponse, err error) {
	if q.loginObject != nil {
		now := time.Now()
		if q.loginObject.ExpiresIn < int(now.Unix()) {
			authRes, err = q.refreshToken()
			if err == nil {
				return
			}
		} else {
			authRes = *q.loginObject
			err = nil
			return
		}
	}
	var requestByte []byte
	var requestBody *bytes.Reader

	requestByte, _ = json.Marshal(map[string]string{
		"terminal_id": q.terminalID,
	})
	requestBody = bytes.NewReader(requestByte)

	url := q.endpoint + QPayAuthToken.Url
	req, err := http.NewRequest(QPayAuthToken.Method, url, requestBody)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(q.username, q.password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return authRes, fmt.Errorf("%s-QPay auth response: %s", time.Now().Format(utils.TimeFormatYYYYMMDDHHMMSS), res.Status)
	}

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &authRes)

	defer res.Body.Close()
	return authRes, nil
}

func (q *qpayquick) refreshToken() (authRes qpayLoginResponse, err error) {
	url := q.endpoint + QPayAuthRefresh.Url
	req, err := http.NewRequest(QPayAuthRefresh.Method, url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+q.loginObject.RefreshToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		return authRes, errors.New(time.Now().Format(utils.TimeFormatYYYYMMDDHHMMSS) + "-QPay token refresh response: " + res.Status)
	}

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &authRes)

	defer res.Body.Close()
	return
}
