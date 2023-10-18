package upoint

import (
	"encoding/json"
)

type upoint struct {
	endpoint   string
	token      string
	terminalID string
}
type Upoint interface {
	CheckUserInfo(input UpointCheckInfoRequest) (UpointCheckInfoResponse, error)
	ProcessTransaction(input UpointTransactionRequest) (response UpointTransactionResponse, err error)
	ReturnTransaction(input UpointReturnTransactionRequest) (response UpointReturnTransactionResponse, err error)
	CheckTransaction(input UpointCheckTransactionRequest) (response UpointCheckTransactionResponse, err error)
	CancelTransaction(input UpointCancelTransactionRequest) (response UpointCancelTransactionResponse, err error)
	GetProducts() (response []UpointProductResponse, err error)
	GetQR() (response UpointQrResponse, err error)
	CheckQR(qr_string string) (response UpointQrCheckResponse, err error)
	CheckQrInfo(qr_string string) (response UpointQrCheckInfoResponse, err error)
	TransactionQR(input UpointTransactionQrRequest) (response UpointTransactionResponse, err error)
}

func New(endpoint, token, terminalID string) Upoint {
	return upoint{
		endpoint:   endpoint,
		token:      token,
		terminalID: terminalID,
	}
}

func (s upoint) CheckUserInfo(input UpointCheckInfoRequest) (response UpointCheckInfoResponse, err error) {
	res, err := s.httpRequestUPoint(input, UpointCheckUserInfo)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) ProcessTransaction(input UpointTransactionRequest) (response UpointTransactionResponse, err error) {
	input.TerminalID = s.terminalID
	res, err := s.httpRequestUPoint(input, UpointProcessTransaction)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) ReturnTransaction(input UpointReturnTransactionRequest) (response UpointReturnTransactionResponse, err error) {
	input.TerminalID = s.terminalID
	res, err := s.httpRequestUPoint(input, UpointReturnTransaction)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) CheckTransaction(input UpointCheckTransactionRequest) (response UpointCheckTransactionResponse, err error) {
	input.TerminalID = s.terminalID
	res, err := s.httpRequestUPoint(input, UpointCheckTransaction)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) CancelTransaction(input UpointCancelTransactionRequest) (response UpointCancelTransactionResponse, err error) {
	res, err := s.httpRequestUPoint(input, UpointCancelTransaction)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) GetProducts() (response []UpointProductResponse, err error) {
	res, err := s.httpRequestUPoint(nil, UpointProduct)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) GetQR() (response UpointQrResponse, err error) {
	res, err := s.httpRequestUPoint(nil, UpointQr)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) CheckQR(qr_string string) (response UpointQrCheckResponse, err error) {
	request := make(map[string]interface{})
	request["qr_string"] = qr_string
	res, err := s.httpRequestUPoint(request, UpointCheckQr)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) CheckQrInfo(qr_string string) (response UpointQrCheckInfoResponse, err error) {
	request := make(map[string]interface{})
	request["qr_string"] = qr_string
	res, err := s.httpRequestUPoint(request, UpointCheckQrInfo)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}

func (s upoint) TransactionQR(input UpointTransactionQrRequest) (response UpointTransactionResponse, err error) {
	input.TerminalID = s.terminalID
	res, err := s.httpRequestUPoint(input, UpointCheckQrInfo)
	if err != nil {
		return
	}
	json.Unmarshal(res, &response)
	return
}
