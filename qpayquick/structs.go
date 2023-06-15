package qpayquick

type (
	qpayLoginResponse struct {
		TokenType        string `json:"token_type"`
		RefreshToken     string `json:"refresh_token"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		Scope            string `json:"scope"`
		NotBeforePolicy  string `json:"not-before-policy"`
		SessionState     string `json:"session_state"`
	}

	QpayCompanyCreateRequest struct {
		OwnerRegNo     string `json:"owner_register_no"`
		OwnerFirstName string `json:"owner_first_name"`
		OwnerLastName  string `json:"owner_last_name"`
		LocationLat    string `json:"location_lat"`
		LocationLng    string `json:"location_lng"`
		RegisterNo     string `json:"register_nubmer"`
		Name           string `json:"name"`
		MCCcode        string `json:"mcc_code"`
		City           string `json:"city"`
		District       string `json:"district"`
		Address        string `json:"address"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
	}
	QpayCompanyCreateResponse struct {
		ID             string `json:"id"`
		VendorID       string `json:"vendor_id"`
		Type           string `json:"type"`
		RegisterNo     string `json:"register_number"`
		Name           string `json:"name"`
		OwnerRegNo     string `json:"owner_register_no"`
		OwnerFirstName string `json:"owner_first_name"`
		OwnerLastName  string `json:"owner_last_name"`
		MCCcode        string `json:"mcc_code"`
		City           string `json:"city"`
		District       string `json:"district"`
		Address        string `json:"address"`
		Phone          string `json:"phone"`
		Email          string `json:"email"`
		LocationLat    string `json:"location_lat"`
		LocationLng    string `json:"location_lng"`
	}

	QpayPersonCreateRequest struct {
		RegisterNo string `json:"register_number"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		MCCcode    string `json:"mcc_code"`
		City       string `json:"city"`
		District   string `json:"district"`
		Address    string `json:"address"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
	}
	QpayPersonCreateResponse struct {
		ID         string `json:"id"`
		VendorID   string `json:"vendor_id"`
		Type       string `json:"type"`
		RegisterNo string `json:"register_number"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		MCCcode    string `json:"mcc_code"`
		City       string `json:"city"`
		District   string `json:"district"`
		Address    string `json:"address"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
	}

	QpayMerchantListRequest struct {
		Offset QpayOffset `json:"offset"`
	}
	QpayMerchantListResponse struct {
		Count int                       `json:"count"`
		Items []QpayMerchantGetResponse `json:"rows"`
	}
	QpayMerchantGetResponse struct {
		CreateDate string `json:"created_date"`
		ID         string `json:"id"`
		Type       string `json:"type"`
		RegisterNo string `json:"register_number"`
		Name       string `json:"name"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		MCCcode    string `json:"mcc_code"`
		City       string `json:"city"`
		District   string `json:"district"`
		Address    string `json:"address"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
	}

	QpayInvoiceRequest struct {
		MerchantID   string                   `json:"merchant_id"`
		BranchCode   string                   `json:"branch_code"`
		Amount       float64                  `json:"amount"`
		Currency     string                   `json:"currency"`
		CustomerName string                   `json:"customer_name"`
		CustomerLogo string                   `json:"customer_logo"`
		CallbackUrl  string                   `json:"callback_url"`
		Description  string                   `json:"description"`
		MCCcode      string                   `json:"mcc_code"`
		BankAccounts []QpayBankAccountRequest `json:"bank_accounts"`
	}

	QpayInvoiceResponse struct {
		ID                  string                    `json:"id"`
		TerminalID          string                    `json:"terminal_id"`
		Amount              string                    `json:"amount"`
		QrCode              string                    `json:"qr_code"`
		Description         string                    `json:"description"`
		InvoiceStatus       string                    `json:"invoice_status"`
		InvoiceStatusDate   string                    `json:"invoice_status_date"`
		CallbackUrl         string                    `json:"callback_url"`
		CustomerName        string                    `json:"customer_name"`
		CustomerLogo        string                    `json:"customer_logo"`
		Currency            string                    `json:"currency"`
		MCCcode             string                    `json:"mcc_code"`
		LegacyID            string                    `json:"legacy_id"`
		VendorID            string                    `json:"vendor_id"`
		ProcessCodeID       string                    `json:"process_code_id"`
		QrImage             string                    `json:"qr_image"`
		InvoiceBankAccounts []QpayBankAccountResponse `json:"invoice_bank_accounts"`
		Urls                []QpayUrls                `json:"urls"`
	}
	QpayUrls struct {
		Name        string `json:"name"`
		Descriptoin string `json:"description"`
		Logo        string `json:"logo"`
		Link        string `json:"link"`
	}
	QpayBankAccountRequest struct {
		AccountBankCode string `json:"account_bank_code"`
		AccountNumber   string `json:"account_number"`
		AccountName     string `json:"account_name"`
		IsDefault       bool   `json:"is_default"`
	}

	QpayInvoiceGetResponse struct {
		ID                  string                    `json:"id"`
		TerminalID          string                    `json:"terminal_id"`
		Amount              string                    `json:"amount"`
		QrCode              string                    `json:"qr_code"`
		Description         string                    `json:"description"`
		InvoiceStatus       string                    `json:"invoice_status"`
		InvoiceStatusDate   string                    `json:"invoice_status_date"`
		CallbackUrl         string                    `json:"callback_url"`
		CustomerName        string                    `json:"customer_name"`
		CustomerLogo        string                    `json:"customer_logo"`
		Currency            string                    `json:"currency"`
		MCCcode             string                    `json:"mcc_code"`
		LegacyID            string                    `json:"legacy_id"`
		VendorID            string                    `json:"vendor_id"`
		ProcessCodeID       string                    `json:"process_code_id"`
		QrImage             string                    `json:"qr_image"`
		InvoiceBankAccounts []QpayBankAccountResponse `json:"invoice_bank_accounts"`
	}

	QpayBankAccountResponse struct {
		ID              string `json:"id"`
		AccountBankCode string `json:"account_bank_code"`
		AccountNumber   string `json:"account_number"`
		AccountName     string `json:"account_name"`
		IsDefault       bool   `json:"is_default"`
		InvoiceID       string `json:"invoice_id"`
	}

	QpayPaymentCheckRequest struct {
		InvoiceID string `json:"invoice_id"`
	}
	QpayPaymentCheckResponse struct {
		ID                string        `json:"id"`
		InvoiceStatus     string        `json:"invoice_status"`
		InvoiceStatusDate string        `json:"invoice_status_date"`
		Payments          []QpayPayment `json:"payments"`
	}

	QpayPayment struct {
		ID                 string             `json:"id"`
		TerminalID         string             `json:"terminal_id"`
		WalletCustomerID   string             `json:"wallet_customer_id"`
		Amount             string             `json:"amount"`
		Currency           string             `json:"currency"`
		PaymentName        string             `json:"payment_name"`
		PaymentDescription string             `json:"payment_description"`
		PaidBy             string             `json:"paid_by"`
		Note               string             `json:"note"`
		PaymentStatus      string             `json:"payment_status"`
		PaymentStatusDate  string             `json:"payment_status_date"`
		Transactions       []QpayTransactions `json:"transactions"`
	}
	QpayTransactions struct {
		ID                  string `json:"id"`
		Description         string `json:"description"`
		TransactionBankCode string `json:"transaction_bank_code"`
		AccountBankCode     string `json:"account_bank_code"`
		AccountBankName     string `json:"account_bank_name"`
		AccountNumber       string `json:"account_number"`
		Status              string `json:"status"`
		Amount              string `json:"amount"`
		Currency            string `json:"currency"`
	}

	QpayOffset struct {
		PageNumber int64 `json:"page_number"` // Хуудасны тоо
		PageLimit  int64 `json:"page_limit"`  // Хуудасны хязгаар max <= 100
	}
)
