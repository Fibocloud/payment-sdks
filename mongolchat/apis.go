package mongolchat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	MchatAuth = utils.API{
		Url:    "/application/auth",
		Method: http.MethodGet,
	}
	MchatOnlineQrGenerate = utils.API{
		Url:    "/worker/onlineqr/generate",
		Method: http.MethodPost,
	}
	MchatOnlineQrcheck = utils.API{
		Url:    "/worker/onlineqr/status",
		Method: http.MethodPost,
	}
	MchatTransactionSettlement = utils.API{
		Url:    "/worker/settle/upload",
		Method: http.MethodPost,
	}
)

func (u *mongolchat) httpRequestMongolChat(body interface{}, api utils.API) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, u.endpoint+api.Url, requestBody)

	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("api-key", u.apikey)
	req.Header.Add("Authorization", "WorkerKey "+u.workerkey)

	res, err := http.DefaultClient.Do(req)
	response, _ = ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = errors.New(string(response))
		fmt.Printf("err here")
		return
	}
	defer res.Body.Close()
	return
}
