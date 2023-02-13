package socialpay

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

// SocialPay
var (
	SocialPayInvoicePhone = utils.API{
		Url:    "/pos/invoice/phone",
		Method: http.MethodPost,
	}
	SocialPayInvoiceQr = utils.API{
		Url:    "/pos/invoice/qr",
		Method: http.MethodPost,
	}
	SocialPayInvoiceCancel = utils.API{
		Url:    "/pos/invoice/cancel",
		Method: http.MethodPost,
	}
	SocialPayInvoiceCheck = utils.API{
		Url:    "/pos/invoice/check",
		Method: http.MethodPost,
	}
	SocialPayPaymentCancel = utils.API{
		Url:    "/pos/payment/cancel",
		Method: http.MethodPost,
	}
	SocialPayPaymentSettlement = utils.API{
		Url:    "/pos/settlement",
		Method: http.MethodPost,
	}
)

func httpRequestSocialpay(body interface{}, api utils.API, endpoint string) (response []byte, err error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, err := http.NewRequest(api.Method, endpoint+api.Url, requestBody)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", utils.HttpContent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = errors.New(string(response))
		fmt.Printf("err here")
		return
	}
	response, _ = io.ReadAll(res.Body)
	defer res.Body.Close()

	return
}
