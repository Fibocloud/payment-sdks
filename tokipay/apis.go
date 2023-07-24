package tokipay

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

// SocialPay
var (
	TokipayPaymentQr = utils.API{
		Url:    "/v4/spose/payment/request",
		Method: http.MethodPost,
	}
	TokipayPaymentSentUser = utils.API{
		Url:    "/v4/spose/payment/user-request",
		Method: http.MethodPost,
	}
	TokipayPaymentScanUser = utils.API{
		Url:    "/v4/spose/payment/scan/user-request",
		Method: http.MethodPost,
	}
	TokipayPaymentStatus = utils.API{
		Url:    "/v4/spose/payment/status?requestId=",
		Method: http.MethodGet,
	}
	TokipayPaymentCancel = utils.API{
		Url:    "/v4/spose/payment/request?requestId=",
		Method: http.MethodDelete,
	}
	TokipayRefund = utils.API{
		Url:    "/v4/spose/payment/refund",
		Method: http.MethodPut,
	}
)

func (q *tokipay) httpRequestTokipay(body interface{}, api utils.API, urlExt string) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, q.endpoint+api.Url+urlExt, requestBody)
	req.Header.Add("Authorization", q.authorization)
	req.Header.Add("api_key", q.apiKey)
	req.Header.Add("im_api_key", q.imApiKey)

	res, err := http.DefaultClient.Do(req)

	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}
	return
}
