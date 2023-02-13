package hipay

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	HipayCheckout = utils.API{
		Url:    "/checkout",
		Method: http.MethodPost,
	}
	HipayCheckoutGet = utils.API{
		Url:    "/checkout/get/",
		Method: http.MethodGet,
	}
	HipayPaymentGet = utils.API{
		Url:    "/payment/get/",
		Method: http.MethodGet,
	}
	HipayPaymentCorrection = utils.API{
		Url:    "/pos/correction",
		Method: http.MethodPost,
	}
	HipayStatement = utils.API{
		Url:    "/pos/statement",
		Method: http.MethodPost,
	}
)

func (h *hipay) httpRequestHipay(body interface{}, api utils.API, urlExt string) (response []byte, err error) {

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, h.endpoint+api.Url+urlExt, requestBody)
	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+h.token)

	res, err := http.DefaultClient.Do(req)

	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}
	return
}
