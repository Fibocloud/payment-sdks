package tokipaythirdparty

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
	TokipayPaymentSentUser = utils.API{
		Url:    "/jump/v1/third-party/payment/request",
		Method: http.MethodPost,
	}
	TokipayPaymentStatus = utils.API{
		Url:    "/jump/v1/third-party/payment/status?requestId=",
		Method: http.MethodGet,
	}
)

func (q *tokipayThirdParty) httpRequestTokipayPOS(body interface{}, api utils.API, urlExt string) (response []byte, err error) {
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
	req.Header.Add("api_key", "third_party_pay")

	res, err := http.DefaultClient.Do(req)

	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}
	return
}
