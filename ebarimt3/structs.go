package ebarimt3

type CreateReceiptInput struct {
	// Required: merchant TIN (ТТД) registered in the PosAPI
	MerchantTin string
	// Required: internal POS number
	PosNo string
	// BranchNo is the seller's branch identifier, e.g. "001"
	BranchNo string
	// BillIDSuffix must be unique within the current day (used to deduplicate ДДТД)
	BillIDSuffix string
	// DistrictCode is the 4-digit region code; defaults to "0101" (UB center)
	DistrictCode string
	// CustomerTin is the buyer's TIN; required for B2B receipts/invoices
	CustomerTin string
	// ConsumerNo is the buyer's 8-digit eBarimt consumer number; used for B2C
	ConsumerNo string
	// Type defaults to B2C_RECEIPT when CustomerTin is empty, B2B_RECEIPT otherwise
	Type ReceiptType
	// TaxType defaults to VAT_ABLE
	TaxType TaxType
	// ClassificationCode is the 7-digit national product/service classification code
	ClassificationCode string
	// Stocks is the list of items/services sold
	Stocks []StockInput
}

// StockInput represents a single sold item or service.
type StockInput struct {
	Code               string
	Name               string
	MeasureUnit        string
	Qty                float64
	UnitPrice          float64 // base price per unit (without taxes)
	CityTax            float64 // city tax amount for this item
	Vat                float64 // VAT amount for this item
	BarCode            string
	BarCodeType        BarCodeType
	ClassificationCode string // overrides CreateReceiptInput.ClassificationCode
	TaxProductCode     string // required when TaxType is VAT_FREE or VAT_ZERO
}

// ─── Wire types sent to the PosAPI ───────────────────────────────────────────

// CreateReceiptRequest is the JSON body for POST /rest/receipt.
type CreateReceiptRequest struct {
	BranchNo     string       `json:"branchNo"`
	TotalAmount  float64      `json:"totalAmount"`
	TotalVAT     *float64     `json:"totalVAT"`
	TotalCityTax *float64     `json:"totalCityTax"`
	DistrictCode string       `json:"districtCode"`
	MerchantTin  string       `json:"merchantTin"`
	PosNo        string       `json:"posNo"`
	CustomerTin  *string      `json:"customerTin"`
	ConsumerNo   *string      `json:"consumerNo"`
	Type         ReceiptType  `json:"type"`
	InactiveId   *string      `json:"inactiveId"`
	ReportMonth  *string      `json:"reportMonth"`
	BillIdSuffix string       `json:"billIdSuffix"`
	InvoiceId    *string      `json:"invoiceId"`
	Receipts     []SubReceipt `json:"receipts"`
	Payments     []Payment    `json:"payments"`
}

// SubReceipt is a sub-receipt inside a batch receipt.
type SubReceipt struct {
	TotalAmount   float64  `json:"totalAmount"`
	TotalVAT      *float64 `json:"totalVAT"`
	TotalCityTax  *float64 `json:"totalCityTax"`
	TaxType       TaxType  `json:"taxType"`
	MerchantTin   string   `json:"merchantTin"`
	CustomerTin   *string  `json:"customerTin,omitempty"`
	BankAccountNo *string  `json:"bankAccountNo"`
	IBan          *string  `json:"iBan"`
	Items         []Item   `json:"items"`
}

// Item is a single product or service line within a sub-receipt.
type Item struct {
	Name               string      `json:"name"`
	BarCode            *string     `json:"barCode"`
	BarCodeType        BarCodeType `json:"barCodeType"`
	ClassificationCode string      `json:"classificationCode"`
	TaxProductCode     *string     `json:"taxProductCode,omitempty"`
	MeasureUnit        string      `json:"measureUnit"`
	Qty                float64     `json:"qty"`
	UnitPrice          float64     `json:"unitPrice"`
	TotalVAT           *float64    `json:"totalVAT"`
	TotalCityTax       *float64    `json:"totalCityTax"`
	TotalAmount        float64     `json:"totalAmount"`
}

// Payment describes how the customer paid.
type Payment struct {
	Code       PaymentCode   `json:"code"`
	Status     PaymentStatus `json:"status"`
	PaidAmount float64       `json:"paidAmount"`
}

// ─── Wire types returned by the PosAPI ───────────────────────────────────────

// CreateReceiptResponse is the response from POST /rest/receipt.
type CreateReceiptResponse struct {
	ID           string  `json:"id"`
	Version      string  `json:"version"`
	TotalAmount  float64 `json:"totalAmount"`
	TotalVAT     float64 `json:"totalVAT"`
	TotalCityTax float64 `json:"totalCityTax"`
	BranchNo     string  `json:"branchNo"`
	DistrictCode string  `json:"districtCode"`
	MerchantTin  string  `json:"merchantTin"`
	PosNo        string  `json:"posNo"`
	Type         string  `json:"type"`

	Receipts []Receipt `json:"receipts"`
	Payments []Payment `json:"payments"`

	PosId   int    `json:"posId"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	QrData  string `json:"qrData"`
	Lottery string `json:"lottery"`
	Date    string `json:"date"`
	Easy    bool   `json:"easy"`
}

type Receipt struct {
	ID           string  `json:"id"`
	TotalAmount  float64 `json:"totalAmount"`
	TaxType      string  `json:"taxType"`
	Items        []Item  `json:"items"`
	MerchantTin  string  `json:"merchantTin"`
	TotalVAT     float64 `json:"totalVAT"`
	TotalCityTax float64 `json:"totalCityTax"`
}

// ReceiptRef is a sub-receipt identifier returned in the response.
type ReceiptRef struct {
	ID            string `json:"id"`
	BankAccountId int    `json:"bankAccountId"`
}

// ReturnReceiptRequest is the body for DELETE /rest/receipt.
type ReturnReceiptRequest struct {
	ID   string `json:"id"`   // 33-digit ДДТД
	Date string `json:"date"` // format: "yyyy-MM-dd HH:mm:ss"
}

// InfoResponse is the response from GET /rest/info.
type InfoResponse struct {
	OperatorName  string `json:"operatorName"`
	OperatorTIN   string `json:"operatorTIN"`
	PosId         int    `json:"posId"`
	PosNo         string `json:"posNo"`
	Version       string `json:"version"`
	LastSentDate  string `json:"lastSentDate"`
	LeftLotteries int    `json:"leftLotteries"`

	AppInfo      AppInfo       `json:"appInfo"`
	PaymentTypes []PaymentType `json:"paymentTypes"`
	Merchants    []Merchant    `json:"merchants"`
	Condition    Condition     `json:"condition"`
}

// AppInfo holds PosAPI application metadata.
type AppInfo struct {
	ApplicationDir    string   `json:"applicationDir"`
	CurrentDir        string   `json:"currentDir"`
	Database          string   `json:"database"`
	DatabaseHost      string   `json:"database-host"`
	SupportedDatabase []string `json:"supported-databases"`
	WorkDir           string   `json:"workDir"`
}
type PaymentType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type Merchant struct {
	TIN       string     `json:"tin"`
	Name      string     `json:"name"`
	VatPayer  bool       `json:"vatPayer"`
	Customers []Customer `json:"customers"`
}
type Customer struct {
	TIN      string `json:"tin"`
	Name     string `json:"name"`
	VatPayer bool   `json:"vatPayer"`
}
type Condition struct {
	IsMedicine bool `json:"isMedicine"`
}