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
		PhoneNumber string `json:"phone_number"`
		Note        string `json:"note"`
	}

	QPaySimpleInvoiceResponse struct {
		InvoiceID    string      `json:"invoice_id"`    // Object id
		QpayShortUrl string      `json:"qPay_shortUrl"` // qr shortcut
		QrText       string      `json:"qr_text"`       // Данс болон картын гүйлгээ дэмжих QR утга
		QrImage      string      `json:"qr_image"`      // Base64  зурган QR код, Qpay лого голдоо агуулсан
		Urls         []*Deeplink `json:"urls"`
	}

	Deeplink struct {
		Name        string `json:"name"`        // Банкны нэр
		Description string `json:"description"` // Утга
		Logo        string `json:"logo"`        // Лого
		Link        string `json:"link"`        // Холбоос линк
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
	// QpayTransaction V2
	QpayTransaction struct {
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
		Count      int64      `json:"count"`       // Нийт гүйлгээний мөрийн тоо
		PaidAmount int64      `json:"paid_amount"` // Гүйлгээний дүн
		Rows       []*QpayRow `json:"rows"`        // Гүйлгээний мөр
	}

	QpayRow struct {
		PaymentID       string      `json:"payment_id"`       // QPay-ээс үүссэн гүйлгээний дугаар
		PaymentStatus   string      `json:"payment_status"`   // Гүйлгээ төлөв // NEW: Гүйлгээ үүсгэгдсэн // FAILED: Гүйлгээ амжилтгүй // PAID: Төлөгдсөн // REFUNDED: Гүйлгээ буцаагдсан
		PaymentDate     interface{} `json:"payment_date"`     // Гүйлгээ хийгдсэн хугацаа
		PaymentFee      string      `json:"payment_fee"`      // Гүйлгээний шимтгэл шимтгэл
		PaymentAmount   string      `json:"payment_amount"`   // Гүйлгээний дүн
		PaymentCurrency string      `json:"payment_currency"` // Гүйлгээний валют
		PaymentWallet   string      `json:"payemnt_wallet"`   // Гүйлгээ хийгдэн воллет
		TransactionType string      `json:"transaction_type"` // P2P or CARD
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
