package qpay_v1

type (
	QpayLogin struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		RefreshToken string `json:"refresh_token"`
	}
	QpayLoginResponse struct {
		TokenType        string `json:"token_type"`
		RefreshToken     string `json:"refresh_token"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		Scope            string `json:"scope"`
		NotBeforePolicy  string `json:"not-before-policy"`
		SessionState     string `json:"session_state"`
	}

	QPayInvoiceCreateRequest struct {
		TemplateID  string              `json:"template_id"`
		MerchantID  string              `json:"merchant_id"`
		BranchID    string              `json:"branch_id"`
		PosID       string              `json:"pos_id"`
		BillNo      string              `json:"bill_no"`
		Date        string              `json:"date"`
		Description string              `json:"description"`
		Amount      float64             `json:"amount"`
		BtukCode    string              `json:"btuk_code"`
		VatFlag     string              `json:"vat_flag"`
		Receiver    QPayInvoiceReceiver `json:"receiver"`
	}

	QPayInvoiceReceiver struct {
		ID          string `json:"id"`
		RegisterNo  string `json:"register_no"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber *int   `json:"phone_number"`
		Note        string `json:"note"`
	}

	QPaySimpleInvoiceResponse struct {
		Name         string     `json:"name"`
		Message      string     `json:"message"`
		PaymentID    int        `json:"payment_id"`
		QPayQrcode   string     `json:"qPay_QRcode"`
		QPayQrimage  string     `json:"qPay_QRimage"`
		CustomerID   string     `json:"customer_id"`
		QPayURL      string     `json:"qPay_url"`
		QPayDeeplink []Deeplink `json:"qPay_deeplink"`
	}

	Deeplink struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
	QpayInvoiceGetResponse struct {
		AllowExceed        bool               `json:"allow_exceed"`
		AllowPartial       bool               `json:"allow_partial"`
		CallbackUrl        string             `json:"callback_url"`
		DiscountAmount     int64              `json:"discount_amount"`
		EnableExpiry       bool               `json:"enable_expiry"`
		ExpiryDate         string             `json:"expiry_date"`
		GrossAmount        int64              `json:"gross_amount"`
		Inputs             []*QpayInput       `json:"inputs"`
		InvoiceDescription string             `json:"invoice_description"`
		InvoiceDueDate     interface{}        `json:"invoice_due_date"`
		InvoiceID          string             `json:"invoice_id"`
		InvoiceStatus      string             `json:"invoice_status"`
		Lines              []*QpayLine        `json:"lines"`
		MaximumAmount      int64              `json:"maximum_amount"`
		MinimumAmount      int64              `json:"minimum_amount"`
		Note               string             `json:"note"`
		SenderBranchCode   string             `json:"sender_branch_code"`
		SenderBranchData   string             `json:"sender_branch_data"`
		SenderInvoiceNo    string             `json:"sender_invoice_no"`
		SurchargeAmount    int64              `json:"surcharge_amount"`
		TaxAmount          int64              `json:"tax_amount"`
		TotalAmount        int64              `json:"total_amount"`
		Transactions       []*QpayTransaction `json:"transactions"`
	}

	QpayInput struct {
	}
	// QLine V2
	QpayLine struct {
		Discounts       []interface{} `json:"discounts"`
		LineDescription string        `json:"line_description"`
		LineQuantity    string        `json:"line_quantity"`
		LineUnitPrice   string        `json:"line_unit_price"`
		Note            string        `json:"note"`
		Surcharges      []interface{} `json:"surcharges"`
		TaxProductCode  interface{}   `json:"tax_product_code"`
		Taxes           []interface{} `json:"taxes"`
	}

	QpayPaymentCheckRequest struct {
		ObjectID   string     `json:"object_id" `  // Обьектын төрөл // INVOICE: Нэхэмжлэх // QR: QR код // ITEM: Бүтээгдэхүүн
		ObjectType string     `json:"object_type"` // Обьектын ID //Обьектын төрөл QR үед QR кодыг ашиглана
		Offset     QpayOffset `json:"offset"`
	}

	QpayOffset struct {
		PageNumber int64 `json:"page_number"` // Хуудасны тоо
		PageLimit  int64 `json:"page_limit"`  // Хуудасны хязгаар max <= 100
	}

	QpayPaymentCheckResponse struct {
		PaymentInfo struct {
			PaymentStatus   string            `json:"payment_status"`
			CustomerID      string            `json:"customer_id"`
			QrAccountID     string            `json:"-"`
			ItemID          string            `json:"-"`
			PaymentAmount   string            `json:"-"`
			TransactionID   string            `json:"-"`
			LastPaymentDate string            `json:"-"`
			Transactions    []QpayTransaction `json:"transactions"`
		} `json:"payment_info"`
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	QpayTransaction struct {
		TransactionBankCode      string  `json:"transaction_bank_code"`
		TransactionNo            int     `json:"transaction_no"`
		TransactionDate          string  `json:"transaction_date"`
		TransactionAmount        float32 `json:"transaction_amount"`
		TransactionCurrencyType  int     `json:"transaction_currency_type"`
		BeneficiaryBankCode      string  `json:"beneficiary_bank_code"`
		BeneficiaryBankName      string  `json:"beneficiary_bank_name"`
		BeneficiaryAccountName   string  `json:"beneficiary_account_name"`
		BeneficiaryAccountNumber string  `json:"beneficiary_account_number"`
	}

	QpayPaymentCancelRequest struct {
		CallbackUrl string `json:"callback_url"` // Төлөгдсөн төлбөрийн ID-г Callback URL-аар авсаны дараа тухайн хэсэгт оруулж өгнө. Жш: https://qpay.mn/payment/result?payment_id=a2ab7ab0-80b0-4045-b79a-3052eda1ca89
		Note        string `json:"note"`         // Тэмдэглэл
	}

	QpayPaymentListRequest struct {
		// ObjectID             string     `json:"object_id"`              // Обьектын ID // Обьектын төрөл INVOICE үед Нэхэмлэхийн код(invoice_code) // Обьектын төрөл QR үед QR кодыг ашиглана
		// ObjectType           string     `json:"object_type"`            // Обьектын төрөл // INVOICE: Нэхэмжлэх // MERCHANT: Байгууллага // QR: QR код
		MerchantID           string     `json:"merchant_id"`            // Байгууллагын дугаар
		MerchantBranchCode   string     `json:"merchant_branch_code"`   // Байгууллагын салбарын код
		MerchantTerminalCode string     `json:"merchant_terminal_code"` // Байгууллагын терминал код
		MerchantStaffCode    string     `json:"merchant_staff_code"`    // Ажилтаны код
		Offset               QpayOffset `json:"offset"`
	}
)
