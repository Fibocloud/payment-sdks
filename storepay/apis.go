package storepay

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

// SocialPay
var (
	StorepayAuth = utils.API{
		Url:    "/oauth/token",
		Method: http.MethodPost,
	}
	StorepayLoan = utils.API{
		Url:    "/merchant/loan",
		Method: http.MethodPost,
	}
	StorepayLoanCheck = utils.API{
		Url:    "/merchant/loan/check/",
		Method: http.MethodGet,
	}
	StorepayUserPossibleAmount = utils.API{
		Url:    "/user/possibleAmount",
		Method: http.MethodPost,
	}
)

func (s *storepay) httpRequest(body interface{}, api utils.API, urlExt string) (response []byte, err error) {
	authObj, authErr := s.authStorepay()
	if authErr != nil {
		err = authErr
		return
	}
	s.loginObject = &authObj

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, s.endpoint+api.Url+urlExt, requestBody)
	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+s.loginObject.AccessToken)

	res, err := http.DefaultClient.Do(req)

	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}

	return
}

// AuthStorepay [Login to storepay]
func (s *storepay) authStorepay() (authRes storepayLoginResponse, err error) {
	if s.loginObject != nil {
		now := time.Now().Unix()
		if now < *s.ExpireIn {
			authRes = *s.loginObject
			err = nil
			return
		}
	}
	url := s.endpoint + StorepayAuth.Url + "?grant_type=password&username=" + s.appUsername + "&password=" + s.appPassword
	req, err := http.NewRequest(StorepayAuth.Method, url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(s.username, s.password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		return authRes, fmt.Errorf("%s-Storepay auth response: %s", time.Now().Format(utils.TimeFormatYYYYMMDDHHMMSS), res.Status)
	}

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &authRes)

	defer res.Body.Close()
	*s.ExpireIn = time.Now().Unix() + int64(authRes.ExpiresIn)
	return authRes, nil
}
