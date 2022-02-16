package golomt

import (
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	ECommerceInvoiceCreate = utils.API{
		Url:    "/api/invoice",
		Method: http.MethodPost,
	}

	ECommerceInquiry = utils.API{
		Url:    "/api/inquiry",
		Method: http.MethodPost,
	}

	ECommercePayByToken = utils.API{
		Url:    "/api/pay",
		Method: http.MethodPost,
	}
)
