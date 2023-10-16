package pass

import (
	"errors"
	"fmt"
)

type BaseError struct {
	Code  string `json:"code"`
	Level string `json:"level"`
	Body  string `json:"body"`
}

func (e *BaseError) toError() error {
	if e == nil {
		return nil
	}

	if e.Code == "" {
		return nil
	}

	return errors.New(fmt.Sprintf("[%s] (%s) %s", e.Code, e.Level, e.Body))
}

// Create

type CreateOrderInput struct {
	EcommerceToken string `json:"ecommerce_token"`
	Amount         int64  `json:"amount"`
	CallbackURL    string `json:"callback_url"`
}

type CreateOrderResponse struct {
	StatusCode string          `json:"status_code"`
	Ret        *RetCreateOrder `json:"ret,omitempty"`
	Msg        *BaseError      `json:"msg,omitempty"`
}

type RetCreateOrder struct {
	Shop     string `json:"shop"`
	Amount   string `json:"amount"`
	OrderID  string `json:"order_id"`
	OrderTTL int64  `json:"order_ttl"`
	DBRefNo  string `json:"db_ref_no"`
}

// Inquery

type OrderInqueryInput struct {
	EcommerceToken string `json:"ecommerce_token"`
	OrderID        string `json:"order_id"`
}

type OrderInqueryResponse struct {
	StatusCode string           `json:"status_code"`
	Ret        *RetInqueryOrder `json:"ret,omitempty"`
	Msg        *BaseError       `json:"msg,omitempty"`
}

type RetInqueryOrder struct {
	RespCode     string        `json:"resp_code"`
	RespMsg      string        `json:"resp_msg"`
	Status       string        `json:"status"`
	Amount       string        `json:"amount"`
	CustomerData *CustomerData `json:"customer_data,omitempty"`
	LoyaltyData  *LoyaltyData  `json:"loyalty_data,omitempty"`
	DBRefNo      string        `json:"db_ref_no"`
	ExtraData    *ExtraData    `json:"extra_data,omitempty"`
	StatusText   string        `json:"status_text"`
}

type CustomerData struct {
	UserID   string `json:"user_id"`
	UniqueID string `json:"unique_id"`
}

type LoyaltyData struct {
	KbStatus              string `json:"kb_status"`
	KbTxnID               string `json:"kb_txn_id"`
	KbCardID              string `json:"kb_card_id"`
	KbUsableLP            string `json:"kb_usable_lp"`
	KbLoyaltyPK           string `json:"kb_loyalty_pk"`
	HasKbLoyalty          string `json:"has_kb_loyalty"`
	KbDescription         string `json:"kb_description"`
	KbLimitValue          string `json:"kb_limit_value"`
	KbLoyaltyType         string `json:"kb_loyalty_type"`
	KbDatesOfWeek         string `json:"kb_dates_of_week"`
	KbNoTxnAmount         string `json:"kb_no_txn_amount"`
	KbYesTxnAmount        string `json:"kb_yes_txn_amount"`
	KbLoyaltyProviderName string `json:"kb_loyalty_provider_name"`
}

type ExtraData struct {
	ID               string `json:"_id"`
	Pan              string `json:"pan"`
	RRN              string `json:"rrn"`
	Amount           string `json:"amount"`
	PosID            string `json:"pos_id"`
	OrderID          string `json:"order_id"`
	RespMsg          string `json:"resp_msg"`
	TraceNo          string `json:"trace_no"`
	DateTime         string `json:"date_time"`
	RespCOde         string `json:"resp_code"`
	TerminalID       string `json:"terminal_id"`
	ApprovalCode     string `json:"approved_code"`
	CurrencyCode     string `json:"currency_code"`
	PaymentRequestID string `json:"payment_request_id"`
}

// Notify

type OrderNotifyInput struct {
	EcommerceToken string `json:"ecommerce_token"`
	OrderID        string `json:"order_id"`
	Phone          string `json:"phone"`
}

type OrderNotifyResponse struct {
	StatusCode string          `json:"status_code"`
	Ret        *RetNotifyOrder `json:"ret,omitempty"`
	Msg        *BaseError      `json:"msg,omitempty"`
}

type RetNotifyOrder struct {
	RespCode string  `json:"resp_code"`
	RespMsg  string  `json:"resp_msg"`
	Success  int64   `json:"success"`
	Data     []Datum `json:"data"`
}

type Datum struct {
	Success   bool   `json:"success"`
	MessageID string `json:"message_id"`
}

// Cancel

type OrderCancelInput struct {
	EcommerceToken string `json:"ecommerce_token"`
	OrderID        string `json:"order_id"`
}

type OrderCancelResponse struct {
	StatusCode string          `json:"status_code"`
	Ret        *RetOrderCancel `json:"ret,omitempty"`
	Msg        *BaseError      `json:"msg,omitempty"`
}

type RetOrderCancel struct {
	RespCode    string      `json:"resp_code"`
	RespMsg     string      `json:"resp_msg"`
	Status      string      `json:"status"`
	Amount      string      `json:"amount"`
	LoyaltyData interface{} `json:"loyalty_data"`
	DBRefNo     string      `json:"db_ref_no"`
}

// Return

type OrderVoidInput struct {
	EcommerceToken string `json:"ecommerce_token"`
	OrderID        string `json:"order_id"`
}

type OrderVoidResponse struct {
	StatusCode string        `json:"status_code"`
	Ret        *RetOrderVoid `json:"ret,omitempty"`
	Msg        *BaseError    `json:"msg,omitempty"`
}

type RetOrderVoid struct {
	RespCode    string      `json:"resp_code"`
	RespMsg     string      `json:"resp_msg"`
	Status      string      `json:"status"`
	Amount      string      `json:"amount"`
	LoyaltyData interface{} `json:"loyalty_data"`
	DBRefNo     string      `json:"db_ref_no"`
	ServiceName string      `json:"service_name"`
	DateTime    string      `json:"date_time"`
	TraceNo     string      `json:"trace_no"`
	Rrn         string      `json:"rrn"`
	TerminalID  string      `json:"terminal_id"`
	MerchantID  string      `json:"merchant_id"`
}

type WebhookCallbackResponse struct {
	OrderID          string            `json:"order_id"`
	PaymentRequestID string            `json:"payment_request_id"`
	PosID            string            `json:"pos_id"`
	Operation        string            `json:"operation"`
	IsSuccess        bool              `json:"is_success"`
	Amount           string            `json:"amount"`
	CreatedTime      string            `json:"created_time"`
	CustomerData     *CustomerData     `json:"customer_data,omitempty"`
	ExtraData        map[string]string `json:"extra_data,omitempty"`
	DBRefNo          string            `json:"db_ref_no"`
}
