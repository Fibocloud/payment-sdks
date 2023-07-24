package tokipay

type (

	// Tokipay Payment Qr Request
	TokipayPaymentQrRequest struct {
		SuccessUrl    string `json:"successUrl"`
		FailureUrl    string `json:"failureUrl"`
		OrderId       string `json:"orderId"`
		MerchantId    string `json:"merchantId"`
		Amount        int64  `json:"amount"`
		Notes         string `json:"notes"`
		Authorization string `json:"authorization"`
	}

	TokipayPaymentQrInput struct {
		OrderId     string `json:"orderId"`
		Amount      int64  `json:"amount"`
		Notes       string `json:"notes"`
		PhoneNo     string `json:"phoneNo"`
		CountryCode string `json:"countryCode"`
		RequestId   string `json:"requestId"`
	}

	// Tokipay Payment Response
	TokipayPaymentResponse struct {
		StatusCode int                        `json:"statusCode"`
		Error      string                     `json:"error"`
		Message    string                     `json:"message"`
		Data       TokipayPaymentDataResponse `json:"data"`
		Type       string                     `json:"type"`
	}
	TokipayPaymentDataResponse struct {
		RequestId string `json:"requestId"`
	}

	TokipayPaymentResponseExt struct {
		StatusCode   int    `json:"statusCode"`
		Error        string `json:"error"`
		Message      string `json:"message"`
		ResponseType string `json:"responseType"`
	}

	// Tokipay Payment Sent user request
	TokipayPaymentSentUserRequest struct {
		SuccessUrl    string `json:"successUrl"`
		FailureUrl    string `json:"failureUrl"`
		OrderId       string `json:"orderId"`
		MerchantId    string `json:"merchantId"`
		Amount        int64  `json:"amount"`
		Notes         string `json:"notes"`
		Authorization string `json:"authorization"`
		PhoneNo       string `json:"phoneNo"`
		CountryCode   string `json:"countryCode"`
	}

	// Tokipay Payment Scan user request
	TokipayPaymentScanUserRequest struct {
		SuccessUrl    string `json:"successUrl"`
		FailureUrl    string `json:"failureUrl"`
		OrderId       string `json:"orderId"`
		MerchantId    string `json:"merchantId"`
		Amount        int64  `json:"amount"`
		Notes         string `json:"notes"`
		Authorization string `json:"authorization"`
		RequestId     string `json:"requestId"`
	}

	TokipayRefundRequest struct {
		RequestId    string `json:"requestId"`
		RefundAmount int64  `json:"refundAmount"`
		MerchantId   string `json:"merchantId"`
	}
	TokipayRefundInput struct {
		RequestId    string `json:"requestId"`
		RefundAmount int64  `json:"refundAmount"`
	}
)
