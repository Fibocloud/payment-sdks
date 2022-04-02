package upoint

type (
	MchatQrGenerateInput struct {
		Amount      int            `json:"amount"`
		PaymentUUID string         `json:"payment_uuid"`
		Products    []MchatProduct `json:"products"`
	}
	MchatWebookInput struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	}
	MchatWebookScanQR struct {
		ReferenceNumber string         `json:"reference_number"`
		WhoPaid         string         `json:"who_paid"`
		UserRefID       string         `json:"user_ref_id"`
		TransactionID   string         `json:"transaction_id"`
		GeneratedQRcode string         `json:"generated_qrcode"`
		Amount          float64        `json:"amount"`
		Date            string         `json:"date"`
		Products        []MchatProduct `json:"products"`
	}
	MchatWebookQuickpay struct {
		ReferenceNumber string         `json:"reference_number"`
		WhoPaid         string         `json:"who_paid"`
		UserRefID       string         `json:"user_ref_id"`
		TransactionID   string         `json:"transaction_id"`
		Amount          float64        `json:"amount"`
		Date            string         `json:"date"`
		Products        []MchatProduct `json:"products"`
	}
	MchatWebookOrder struct {
		ReferenceNumber string         `json:"reference_number"`
		WhoPaid         string         `json:"who_paid"`
		UserRefID       string         `json:"user_ref_id"`
		TransactionID   string         `json:"transaction_id"`
		OrderID         string         `json:"order_id"`
		Amount          float64        `json:"amount"`
		Date            string         `json:"date"`
		Products        []MchatProduct `json:"products"`
	}
	MchatProduct struct {
		ProductName string `json:"product_name"`
		Quantity    string `json:"quantity"`
		Price       int    `json:"price"`
		Tag         string `json:"tag"`
	}
	MchatOnlineQrGenerateRequest struct {
		Amount          int            `json:"amount"`
		BranchID        string         `json:"branch_id"` // optional
		Products        []MchatProduct `json:"products"`
		Title           string         `json:"title"`
		SubTitle        string         `json:"sub_title"`
		NOAT            string         `json:"noat"`
		NHAT            string         `json:"nhat"`
		TTD             string         `json:"ttd"`
		ReferenceNumber string         `json:"reference_number"`
		ExpireTime      string         `json:"expire_time"`
	}
	MchatOnlineQrGenerateResponse struct {
		Qr      string `json:"qr"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	MchatOnlineQrCheckResponse struct {
		Status    string `json:"status"`
		Code      int    `json:"code"`
		Message   string `json:"message"`
		ID        string `json:"id"`          // if status is paid this field filled with transaction id
		WhoPaid   string `json:"who_paid"`    // if status is paid this field filled with who_paid
		UserRefID string `json:"user_ref_id"` // if status is paid this field filled with user_ref_id
	}
)
