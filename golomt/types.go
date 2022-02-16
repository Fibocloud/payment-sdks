package golomt

type Lang string

const (
	MN = Lang("MN")
	EN = Lang("EN")
)

type PaymentMethod string

const (
	EcommercePay = PaymentMethod("payment")
	SocialPay    = PaymentMethod("socialpay")
)

type ReturnType string

const (
	POST   = ReturnType("POST")
	GET    = ReturnType("GET")
	MOBILE = ReturnType("MOBILE")
)
