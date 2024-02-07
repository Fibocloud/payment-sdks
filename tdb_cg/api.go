package tdb_cg

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

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
	QPayPaymentGet = utils.API{
		Url:    "/payment/get/",
		Method: http.MethodGet,
	}
	QPayPaymentCheck = utils.API{
		Url:    "/payment/check/",
		Method: http.MethodGet,
	}

	QPayInvoiceCreate = utils.API{
		Url:    "/bill/create",
		Method: http.MethodPost,
	}
	QPayInvoiceGet = utils.API{
		Url:    "/invoice/",
		Method: http.MethodGet,
	}
)

// func (q *qpay) ExpireTokenForce() {
// 	q.loginObject.ExpiresIn = 0
// }

func (q *tdbcg) httpRequest(body interface{}, api utils.API, urlExt string) (response []byte, err error) {

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	// pfxData, err := os.ReadFile(q.certPathPfx)
	// if err != nil {
	// 	fmt.Println("Error reading PFX file:", err)
	// 	return nil, err
	// }
	// cerData, err := os.ReadFile(q.certPathCer)
	// if err != nil {
	// 	fmt.Println("Error reading PFX file:", err)
	// 	return nil, err
	// }

	// _, certification, err := pkcs12.Decode(pfxData, q.certPass)
	// derBytes, err := x509.MarshalCertificate(certification)
	// pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	cert, err := tls.LoadX509KeyPair(q.certPathCer, "keyandcerts.pem")
	if err != nil {
		fmt.Println("Error loading certificate:", err)
		return
	}

	// Create a TLS configuration with the loaded certificate
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	req, _ := http.NewRequest(api.Method, q.endpoint+api.Url+urlExt, requestBody)
	req.Header.Add("Content-Type", utils.XmlContent)
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(q.loginID+":"+q.pwd)))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()
	response, _ = io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(response))
	}

	return
}
