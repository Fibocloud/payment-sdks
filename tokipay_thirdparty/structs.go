package tokipaythirdparty

type (
	TokipayPaymentInput struct {
		OrderId     string `json:"orderId"`
		Amount      int64  `json:"amount"`
		Notes       string `json:"notes"`
		PhoneNo     string `json:"phoneNo"`
		CountryCode string `json:"countryCode"`
		RequestId   string `json:"requestId"`
		SuccessUrl  string `json:"successUrl"`
	}

	// Tokipay Payment Response
	TokipayPaymentResponse struct {
		StatusCode int                           `json:"statusCode"`
		Error      string                        `json:"error"`
		Message    string                        `json:"message"`
		Data       TokipayPaymentRequestResponse `json:"data"`
		Type       string                        `json:"type"`
	}

	TokipayPaymentRequestResponse struct {
		RequestId string `json:"requestId"`
	}

	TokipayPaymentStatusResponse struct {
		StatusCode int                              `json:"statusCode"`
		Error      string                           `json:"error"`
		Message    string                           `json:"message"`
		Data       TokipayPaymentStatusDataResponse `json:"data"`
		Type       string                           `json:"type"`
	}
	TokipayPaymentStatusDataResponse struct {
		Status string `json:"status"`
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
)
