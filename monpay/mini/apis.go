package mini

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	MiniAppGetAccessToken = utils.API{
		Url:    "/oauth/token",
		Method: http.MethodPost,
	}
	MiniAppGetUserInfo = utils.API{
		Url:    "/api/oauth/userinfo",
		Method: http.MethodGet,
	}
	MiniAppCreateInvoice = utils.API{
		Url:    "/api/oauth/invoice",
		Method: http.MethodPost,
	}
)

type MiniAppInput struct {
	ClientId     string
	ClientSecret string
	Code         string
	GrantType    string
	RedirectUri  string
}

type miniAppLoginReponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func (m *mini) getAccessToken() error {
	ep := m.endpoint + MiniAppGetAccessToken.Url
	data := url.Values{}
	data.Set("redirect_uri", m.redirecturi)
	data.Set("client_id", m.clientsecret)
	data.Set("client_secret", m.clientsecret)
	data.Set("code", m.code)
	data.Set("grant_type", m.granttype)
	req, err := http.NewRequest(MiniAppGetAccessToken.Method, ep, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Content-Type", "x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	var response miniAppLoginReponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	m.accesstoken = response.AccessToken
	m.expiresin = response.ExpiresIn
	m.refreshtoken = response.RefreshToken
	return nil
}

func (m *mini) miniApphttpRequest(body interface{}, api utils.API, urlExt string, data interface{}) (err error) {
	if m.accesstoken == "" || time.Now().Unix() > int64(m.expiresin) {
		if err = m.getAccessToken(); err != nil {
			return
		}
	}
	ep := m.endpoint + api.Url
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ := json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}
	req, err := http.NewRequest(api.Method, ep, requestBody)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer auth "+m.accesstoken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	response, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = errors.New(string(response))
		return
	}

	var resp MiniAppResponse

	err = json.Unmarshal(response, &resp)
	if err != nil {
		return
	}

	if resp.IntCode != 200 {
		err = errors.New(resp.Info)
		return
	}

	result, err := json.Marshal(resp.Result)
	if err != nil {
		return
	}

	err = json.Unmarshal(result, data)
	return
}
