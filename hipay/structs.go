package hipay

type (
	HipayCheckoutRequest struct {
		EntityID string          `json:"entityId"`
		Amount   float64         `json:"amount"`
		Currency string          `json:"currency"`
		QrData   bool            `json:"qrData"`
		Signal   bool            `json:"signal"`
		IpAddres string          `json:"ipaddress"`
		Items    []*CheckOutItem `json:"items,omitempty"`
	}
	CheckOutItem struct {
		ItemNo   string  `json:"itemno"`
		Name     string  `json:"names"`
		Price    float64 `json:"price"`
		Quantity int64   `json:"quantity"`
		Brand    string  `json:"brand"`
		Measure  string  `json:"measure"`
		Vat      float64 `json:"vat"`
		CityTax  float64 `json:"citytax"`
	}
	HipayCheckoutResponse struct {
		Code        int64      `json:"code"`
		Description string     `json:"description"`
		RequestID   string     `json:"requestId"`
		CheckoutID  string     `json:"checkoutId"`
		Expires     string     `json:"expires"`
		Signal      *Signal    `json:"signal,omitempty"`
		QrData      string     `json:"qrData"`
		Details     []*Details `json:"details,omitempty"`
	}
	Signal struct {
		SubscribeKey string `json:"subscribeKey"`
		Channel      string `json:"channel"`
		UUID         string `json:"uuid"`
	}

	HipayCheckoutGetResponse struct {
		Code           int64      `json:"code"`
		Description    string     `json:"description"`
		Amount         string     `json:"amount,omitempty"`
		Currency       string     `json:"currency,omitempty"`
		DiscountAmount string     `json:"discount_amount,omitempty"`
		Status         string     `json:"status,omitempty"`
		StatusDate     string     `json:"status_date,omitempty"`
		PaymentID      string     `json:"paymentId,omitempty"`
		Details        []*Details `json:"details,omitempty"`
	}
	Details struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
	}

	HipayPaymentGetResponse struct {
		Code         int64      `json:"code"`
		Description  string     `json:"description"`
		ID           string     `json:"id,omitempty"`
		Amount       string     `json:"amount,omitempty"`
		Currency     string     `json:"currency,omitempty"`
		EntityID     string     `json:"entityId,omitempty"`
		CheckoutID   string     `json:"checkoutId,omitempty"`
		PaymentID    string     `json:"paymentId,omitempty"`
		PaymentType  string     `json:"paymentType,omitempty"`
		PaymentBrand string     `json:"paymentBrand,omitempty"`
		PaymentDate  string     `json:"paymentDate,omitempty"`
		PaymentDesc  string     `json:"paymentDesc,omitempty"`
		ResultDesc   string     `json:"result_desc,omitempty"`
		ResultCode   string     `json:"result_code,omitempty"`
		Details      []*Details `json:"details,omitempty"`
	}

	HipayPaymentCorrectionRequest struct {
		EntityID  string `json:"entityId"`
		PaymentID string `json:"paymentId"`
	}

	HipayPaymentCorrectionResponse struct {
		Code                int64      `json:"code"`
		Description         string     `json:"description"`
		PaymentID           string     `json:"paymentId,omitempty"`
		CorrectionPaymentID string     `json:"correction_paymentId,omitempty"`
		Details             []*Details `json:"details,omitempty"`
	}

	HipayStatementRequest struct {
		EntityID string `json:"entityId"`
		Date     string `json:"date"` // Format: "2023-02-09"
	}
	HipayStatementResponse struct {
		Code        int64      `json:"code"`
		Description string     `json:"description"`
		Details     []*Details `json:"details,omitempty"`
		Data        *Data      `json:"data,omitempty"`
	}

	Data struct {
		List       []*List `json:"list"`
		EntityID   string  `json:"entityId"`
		Date       string  `json:"date"`
		Page       int64   `json:"page"`
		PerPage    int64   `json:"perPage"`
		TotalCount int64   `json:"totalCount"`
		TotalPage  int64   `json:"totalPage"`
	}
	List struct {
		PaymentDate  string  `json:"paymentDate"`
		CheckoutID   string  `json:"checkoutId"`
		PaymentID    string  `json:"paymentId"`
		Amount       float64 `json:"amount"`
		Currency     string  `json:"currency"`
		FeeAmount    float64 `json:"feeAmount"`
		FeePrc       float64 `json:"feePrc"`
		PaymentDesc  string  `json:"paymentDesc"`
		ReturnAmount float64 `json:"returnAmount"`
	}
)
