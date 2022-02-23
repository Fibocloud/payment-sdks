package mini

import "time"

type (
	MiniAppResponse struct {
		Code    string                 `json:"code"`
		IntCode int                    `json:"intCode"`
		Info    string                 `json:"info"`
		Result  map[string]interface{} `json:"result"`
	}

	MiniAppUserInfo struct {
		UserId        int    `json:"userId"`
		UserPhone     string `json:"userPhone"`
		UserEmail     string `json:"userEmail"`
		UserFirstname string `json:"userFirstname"`
		UserLastname  string `json:"userLastname"`
	}

	CreateInvoiceRequest struct {
		Amount           float64 `json:"amount"`
		RedirectUri      string  `json:"redirectUri"`
		ClientServiceUrl string  `json:"clientServiceUrl"`
		Receiver         string  `json:"receiver"`
		InvoiceType      string  `json:"invoiceType"`
		Description      string  `json:"description"`
	}

	CreateInvoiceResponse struct {
		Id          string    `json:"id"`
		Amount      float64   `json:"amount"`
		Status      string    `json:"status"`
		InvoiceType string    `json:"invoiceType"`
		Receiver    string    `json:"receiver"`
		Description string    `json:"description"`
		CreateDate  time.Time `json:"createDate"`
		UpdateDate  time.Time `json:"updateDate"`
		MiniAppId   int       `json:"miniAppId"`
		UserId      int       `json:"userId"`
		RedirectUri string    `json:"redirectUri"`
	}
)
