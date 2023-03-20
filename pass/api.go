package pass

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	PassCreateOrder = utils.API{
		Url:    "/create_order",
		Method: http.MethodPost,
	}
	PassInqueryOrder = utils.API{
		Url:    "/order_inquiry",
		Method: http.MethodPost,
	}
	PassNotifyOrder = utils.API{
		Url:    "/order_notify",
		Method: http.MethodPost,
	}
	PassCancelOrder = utils.API{
		Url:    "/cancel_order",
		Method: http.MethodPost,
	}
	PassVoidOrder = utils.API{
		Url:    "/void",
		Method: http.MethodPost,
	}
)

func (p *pass) httpRequestPass(body interface{}, api utils.API, urlExt string) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, p.endpoint+api.Url+urlExt, requestBody)
	req.Header.Add("Content-Type", utils.HttpContent)

	res, err := http.DefaultClient.Do(req)

	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}
	return
}
