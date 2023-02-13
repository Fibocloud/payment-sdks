package golomt

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

type golomt struct {
	endpoint    string
	secret      string
	bearerToken string
}

type Ecommerce interface {
	CreateInvoice(input CreateInvoiceInput) (CreateInvoiceResponse, error)
	Inquiry(transactionId string) (InquiryResponse, error)
	PayByToken(amount float64, token string, transactionId string, returnUrl string) (ByTokenResponse, error)
	GetUrlByInvoiceId(invoice string, lang Lang, paymentMethod PaymentMethod) string
}

func New(endpoint, secret, token string) Ecommerce {
	return golomt{
		endpoint:    endpoint,
		secret:      secret,
		bearerToken: token,
	}
}

func boolToString(v bool) string {
	if v {
		return "Y"
	}
	return "N"
}

func (g golomt) GetUrlByInvoiceId(invoice string, lang Lang, paymentMethod PaymentMethod) string {
	return fmt.Sprintf(g.endpoint+"/%v/%v/%v", paymentMethod, lang, invoice)
}

func (g golomt) PayByToken(amount float64, token string, transactionId string, lang string) (response ByTokenResponse, err error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(amount, transactionId, token))
	request := ByTokenRequest{
		Amount:        fmt.Sprintf("%v", amount),
		Checksum:      checksum,
		Token:         token,
		TransactionID: transactionId,
		Lang:          lang,
	}

	res, err := g.httpRequestGolomtEcommerce(request, ECommercePayByToken, g.endpoint)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, &response); err != nil {
		return
	}

	if response.ErrorCode != "000" {
		err = fmt.Errorf("%s", response.ErrorDesc)
		return
	}
	return
}

func (g golomt) CreateInvoice(input CreateInvoiceInput) (response CreateInvoiceResponse, err error) {
	fmt.Printf("golomt cred: %v\n", g)
	fmt.Printf("golomt input: %v\n", input)
	_amount := fmt.Sprintf("%.2f", input.Amount)

	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		input.TransactionID,
		_amount,
		input.ReturnType,
		input.Callback,
	))

	request := CreateInvoiceRequest{
		Amount:        _amount,
		Checksum:      checksum,
		GenerateToken: boolToString(input.GetToken),
		Callback:      input.Callback,
		TransactionID: input.TransactionID,
		ReturnType:    string(input.ReturnType),
	}

	fmt.Printf("checksum: %s\n", checksum)
	fmt.Printf("golomt request: %v\n", request)

	res, err := g.httpRequestGolomtEcommerce(request, ECommerceInvoiceCreate, g.endpoint)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Printf("golomt response: %v\n", string(res))

	if err = json.Unmarshal(res, &response); err != nil {
		fmt.Println("Unmarshal: ", err)
		return
	}

	return
}

func (g golomt) Inquiry(transactionId string) (response InquiryResponse, err error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(transactionId, transactionId))
	request := InquiryRequest{
		Checksum:      checksum,
		TransactionID: transactionId,
	}

	res, err := g.httpRequestGolomtEcommerce(request, ECommerceInquiry, g.endpoint)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, &response); err != nil {
		return
	}

	if response.ErrorCode != "000" {
		err = fmt.Errorf("%s", response.ErrorDesc)
		return
	}
	return
}

func (g golomt) httpRequestGolomtEcommerce(body interface{}, api utils.API, endpoint string) (response []byte, err error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, endpoint+api.Url, requestBody)
	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+g.bearerToken)

	res, err := http.DefaultClient.Do(req)
	response, err = io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var error Error
		if err := json.Unmarshal(response, &error); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", error.Message)
	}
	return
}
