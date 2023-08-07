package storepay

type (
	storepayLoginResponse struct {
		TokenType      string     `json:"token_type"`
		RefreshToken   string     `json:"refresh_token"`
		AccessToken    string     `json:"access_token"`
		ExpiresIn      int        `json:"expires_in"`
		Scope          string     `json:"scope"`
		CurrentStoreId int        `json:"current_store_id"`
		UserId         int        `json:"user_id"`
		RoleId         RoleStruct `json:"role_id"`
		Jti            string     `json:"jti"`
	}
	RoleStruct struct {
		Id          int    `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`
		Authority   string `json:"authority"`
	}

	StorepayLoanInput struct {
		Description  string  `json:"description"`
		MobileNumber string  `json:"mobileNumber"`
		Amount       float64 `json:"amount"`
	}

	StorepayLoanRequest struct {
		StoreId      string `json:"storeId"`
		MobileNumber string `json:"mobileNumber"`
		Description  string `json:"description"`
		Amount       string `json:"amount"`
		CallbackUrl  string `json:"callbackUrl"`
	}
	StorepayLoanResponse struct {
		Value   string      `json:"value"`
		MsgList []MsgStruct `json:"msgList"`
		Attrs   interface{} `json:"attrs"`
		Status  string      `json:"status"`
	}
	StorepayCheckResponse struct {
		Value   bool        `json:"value"`
		MsgList []MsgStruct `json:"msgList"`
		Attrs   interface{} `json:"attrs"`
		Status  string      `json:"status"`
	}
	MsgStruct struct {
		Code   string `json:"code"`
		Text   string `json:"text"`
		Params string `json:"params"`
	}

	StorepayUserCheckRequest struct {
		MobileNumber string `json:"mobileNumber"`
	}
	StorepayUserCheckResponse struct {
		Value   float64     `json:"value"`
		MsgList []MsgStruct `json:"msgList"`
		Attrs   interface{} `json:"attrs"`
		Status  string      `json:"status"`
	}
)
