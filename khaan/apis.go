package khaan

import (
	"net/http"

	"github.com/Fibocloud/payment-sdks/utils"
)

var (
	KhaanOrderRegister = utils.API{
		Url:    "/register.do",
		Method: http.MethodPost,
	}

	KhaanOrderStatus = utils.API{
		Url:    "/getOrderStatus.do",
		Method: http.MethodPost,
	}
)
