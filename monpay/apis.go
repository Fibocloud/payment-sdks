package monpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	MonpayGenerateQr = utils.API{
		Url:    "/rest/branch/qrpurchase/generate",
		Method: http.MethodPost,
	}
	MonpayCheckQr = utils.API{
		Url:    "/rest/branch/qrpurchase/check?uuid=",
		Method: http.MethodGet,
	}
)

func (m *monpay) httpRequestMonpay(body interface{}, api utils.API, urlExt string) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, m.endpoint+api.Url+urlExt, requestBody)
	req.SetBasicAuth(m.username, m.accoutnId)
	req.Header.Add("Content-Type", utils.HttpContent)

	res, err := http.DefaultClient.Do(req)
	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = errors.New(string(response))
		fmt.Print(err.Error())
		return
	}
	defer res.Body.Close()
	return
}
