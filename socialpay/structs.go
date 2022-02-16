package socialpay

type (
	SocialPayInvoicePhoneRequest struct {
		Phone    string `json:"phone"`
		Amount   string `json:"amount"`
		Invoice  string `json:"invoice"`
		Terminal string `json:"terminal"`
		Checksum string `json:"checksum"`
	}

	SocialPayInvoiceSimpleRequest struct {
		Amount   string `json:"amount"`
		Invoice  string `json:"invoice"`
		Terminal string `json:"terminal"`
		Checksum string `json:"checksum"`
	}

	SocialPaySettlementRequest struct {
		SettlementId string `json:"settlementId"`
		Checksum     string `json:"checksum"`
		Terminal     string `json:"terminal"`
	}

	SocialPaySimpleResponse struct {
		Description string `json:"desc"`
		Status      string `json:"status"`
	}

	SocialPayTransactionResponse struct {
		ApprovalCode        string  `json:"approval_code"`
		Amount              float64 `json:"amount"`
		CardNumber          string  `json:"card_number"`
		ResponseDescription string  `json:"resp_desc"`
		ResponseCode        string  `json:"resp_code"`
		Terminal            string  `json:"terminal"`
		Invoice             string  `json:"invoice"`
		Checksum            string  `json:"checksum"`
	}

	SocialPayPaymentSettlementResponse struct {
		Amount float64 `json:"amount"`
		Count  int     `json:"count"`
		Status string  `json:"status"`
	}

	SocialPayErrorResponse struct {
		ErrorDescription string `json:"errorDesc"`
		ErrorType        string `json:"errorType"`
	}

	Header struct {
		Status string `json:"status"`
		Code   int    `json:"code"`
	}

	Body struct {
		Response map[string]interface{} `json:"response"`
		Error    map[string]interface{} `json:"error"`
	}

	Response struct {
		Header Header `json:"header"`
		Body   Body   `json:"body"`
	}
)
