package monpay

type (
	MonpayQrInput struct {
		Amount float64
	}
	MonpayQrRequest struct {
		Amount       float64 `json:"amount"`
		DisplayName  string  `json:"string"`
		GenerateUUID bool    `json:"generateUuid"`
		CallbackUrl  string  `json:"callbackUrl"`
	}
	MonpayQrResponse struct {
		Code   int            `json:"code"`
		Info   string         `json:"info"`
		Result MonpayResultQr `json:"result"`
	}
	MonpayResultQr struct {
		Qrcode string `json:"qrcode"`
		UUID   string `json:"uuid"`
	}

	MonpayCheckResponse struct {
		Code   int               `json:"code"`
		Info   string            `json:"info"`
		Result MonpayResultCheck `json:"result"`
	}
	MonpayResultCheck struct {
		UUID          string `json:"uuid"`
		UsedAt        int64  `json:"usedAt"`
		UsedByUd      int64  `json:"usedById"`
		TransactionId string `json:"transactionId"`
		Amount        int64  `json:"amount"`
		CreatedAt     int64  `json:"createdAt"`
		UserPhone     string `json:"userPhone"`
		UserAccountNo string `json:"userAccountNo"`
		UserVatId     string `json:"userVatId"`
		UsedAtUI      string `json:"usedAtUI"`
		CreatedAtUI   string `json:"createdAtUI"`
		AmountUI      string `json:"amountUI"`
	}

	MonpayCallback struct {
		Amount float64 `schema:"amount"`
		UUID   string  `schema:"uuid"`
		Status string  `schema:"status"`
		TnxId  string  `schema:"tnxId"`
	}
)
