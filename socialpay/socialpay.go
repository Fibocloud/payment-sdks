package socialpay

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Fibocloud/payment-sdks/utils"
)

type socialPay struct {
	terminal string
	secret   string
	endpoint string
}
type SocialPay interface {
	InvoicePhone(amount float64, invoice, phone string) (SocialPaySimpleResponse, error)
	InvoiceQR(amount float64, invoice string) (SocialPaySimpleResponse, error)
	CheckTransaction(amount float64, invoice string) (SocialPayTransactionResponse, error)
	CancelInvoice(amount float64, invoice string) (SocialPaySimpleResponse, error)
	CancelTransaction(amount float64, invoice string) (SocialPayTransactionResponse, error)
	TransactionSettlement(settlementId string) (SocialPayPaymentSettlementResponse, error)
}

func New(terminal, secret, endpoint string) SocialPay {
	return socialPay{
		terminal: terminal,
		secret:   secret,
		endpoint: endpoint,
	}
}

func (s socialPay) TransactionSettlement(settlementId string) (response SocialPayPaymentSettlementResponse, err error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, settlementId))
	request := SocialPaySettlementRequest{
		SettlementId: settlementId,
		Checksum:     checksum,
		Terminal:     s.terminal,
	}

	res, err := httpRequestSocialpay(request, SocialPayPaymentSettlement, s.endpoint)
	if err != nil {
		return
	}
	var resp Response
	json.Unmarshal(res, &resp)

	if resp.Header.Code != 200 {
		errorResponse := mapToErrorResponse(resp.Body.Error)
		err = errors.New(errorResponse.ErrorDescription)
	} else {
		response = mapToSettlementResponse(resp.Body.Response)
	}
	return
}

func (s socialPay) CancelTransaction(amount float64, invoice string) (response SocialPayTransactionResponse, err error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, invoice, amount))
	request := SocialPayInvoiceSimpleRequest{
		Amount:   fmt.Sprintf("%v", amount),
		Invoice:  invoice,
		Terminal: s.terminal,
		Checksum: checksum,
	}

	res, err := httpRequestSocialpay(request, SocialPayPaymentCancel, s.endpoint)
	if err != nil {
		return
	}
	var resp Response
	json.Unmarshal(res, &resp)

	if resp.Header.Code != 200 {
		errorResponse := mapToErrorResponse(resp.Body.Error)
		err = errors.New(errorResponse.ErrorDescription)
	} else {
		response = mapToTransactionInfo(resp.Body.Response)
	}
	return
}

func (s socialPay) CheckTransaction(amount float64, invoice string) (response SocialPayTransactionResponse, err error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, invoice, amount))
	request := SocialPayInvoiceSimpleRequest{
		Amount:   fmt.Sprintf("%v", amount),
		Invoice:  invoice,
		Terminal: s.terminal,
		Checksum: checksum,
	}
	res, err := httpRequestSocialpay(request, SocialPayInvoiceCheck, s.endpoint)
	if err != nil {
		return
	}
	var resp Response
	json.Unmarshal(res, &resp)

	if resp.Header.Code != 200 {
		errorResponse := mapToErrorResponse(resp.Body.Error)
		err = errors.New(errorResponse.ErrorDescription)
	} else {
		response = mapToTransactionInfo(resp.Body.Response)
	}
	return
}

func (s socialPay) InvoicePhone(amount float64, invoice, phone string) (response SocialPaySimpleResponse, err error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, invoice, amount, phone))
	request := SocialPayInvoicePhoneRequest{
		Phone:    phone,
		Amount:   fmt.Sprintf("%v", amount),
		Invoice:  invoice,
		Terminal: s.terminal,
		Checksum: checksum,
	}
	res, err := httpRequestSocialpay(request, SocialPayInvoicePhone, s.endpoint)
	if err != nil {
		return
	}
	var resp Response
	json.Unmarshal(res, &resp)

	if resp.Header.Code != 200 {
		errorResponse := mapToErrorResponse(resp.Body.Error)
		err = errors.New(errorResponse.ErrorDescription)
	} else {
		response = mapToSimpleResponse(resp.Body.Response)
	}
	return
}

func (s socialPay) InvoiceQR(amount float64, invoice string) (response SocialPaySimpleResponse, err error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, invoice, amount))
	request := SocialPayInvoiceSimpleRequest{
		Amount:   fmt.Sprintf("%v", amount),
		Invoice:  invoice,
		Terminal: s.terminal,
		Checksum: checksum,
	}
	res, err := httpRequestSocialpay(request, SocialPayInvoiceQr, s.endpoint)
	if err != nil {
		return
	}
	var resp Response
	json.Unmarshal(res, &resp)

	if resp.Header.Code != 200 {
		errorResponse := mapToErrorResponse(resp.Body.Error)
		err = errors.New(errorResponse.ErrorDescription)
	} else {
		response = mapToSimpleResponse(resp.Body.Response)
	}
	return
}

func (s socialPay) CancelInvoice(amount float64, invoice string) (response SocialPaySimpleResponse, err error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, invoice, amount))
	request := SocialPayInvoiceSimpleRequest{
		Amount:   fmt.Sprintf("%v", amount),
		Invoice:  invoice,
		Terminal: s.terminal,
		Checksum: checksum,
	}
	res, err := httpRequestSocialpay(request, SocialPayInvoiceCancel, s.endpoint)
	if err != nil {
		return
	}
	var resp Response
	json.Unmarshal(res, &resp)

	if resp.Header.Code != 200 {
		errorResponse := mapToErrorResponse(resp.Body.Error)
		err = errors.New(errorResponse.ErrorDescription)
	} else {
		response = mapToSimpleResponse(resp.Body.Response)
	}
	return
}
