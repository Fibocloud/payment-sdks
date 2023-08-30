package golomt

type (
	CreateInvoiceInput struct {
		Amount         float64    `json:"amount"`
		TransactionID  string     `json:"transactionId"`
		ReturnType     ReturnType `json:"returnType"`
		Callback       string     `json:"callback"`
		GetToken       bool       `json:"getToken"`
		SocialDeeplink bool       `json:"socialDeeplink"`
	}

	CreateInvoiceResponse struct {
		Invoice        string `json:"invoice"`
		Checksum       string `json:"checksum"`
		TransactionID  string `json:"transactionId"`
		Timestamp      string `json:"timestamp"`
		Status         int    `json:"status"`
		Error          string `json:"error"`
		Message        string `json:"message"`
		Path           string `json:"path"`
		SocialDeeplink string `json:"socialDeeplink"`
	}

	CreateInvoiceRequest struct {
		Amount         string `json:"amount"`
		Checksum       string `json:"checksum"`
		TransactionID  string `json:"transactionId"`
		ReturnType     string `json:"returnType"`
		Callback       string `json:"callback"`
		GenerateToken  string `json:"genToken"`
		SocialDeeplink string `json:"socialDeeplink"`
	}

	InquiryResponse struct {
		Amount        string `json:"amount"`
		Bank          string `json:"bank"`
		Status        string `json:"status"`
		ErrorDesc     string `json:"errorDesc"`
		ErrorCode     string `json:"errorCode"`
		CardHolder    string `json:"cardHolder"`
		CardNumber    string `json:"cardNumber"`
		TransactionID string `json:"transactionId"`
		Token         string `json:"token"`
	}

	InquiryRequest struct {
		Checksum      string `json:"checksum"`
		TransactionID string `json:"transactionId"`
	}

	ByTokenRequest struct {
		Amount        string `json:"amount"`
		Invoice       string `json:"invoice"`
		Checksum      string `json:"checksum"`
		TransactionID string `json:"transactionId"`
		Token         string `json:"token"`
		Lang          string `json:"lang"`
	}

	ByTokenResponse struct {
		Amount        string `json:"amount"`
		ErrorDesc     string `json:"errorDesc"`
		ErrorCode     string `json:"errorCode"`
		TransactionID string `json:"transactionId"`
		Checksum      string `json:"checksum"`
		CardNumber    string `json:"cardNumber"`
	}

	Error struct {
		Timestamp string `json:"timestamp"`
		Status    int    `json:"status"`
		Error     string `json:"error"`
		Message   string `json:"message"`
		Path      string `json:"path"`
	}
)
