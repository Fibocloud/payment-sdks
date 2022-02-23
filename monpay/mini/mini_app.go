package mini

type mini struct {
	accesstoken  string
	expiresin    int
	refreshtoken string
	redirecturi  string
	clientid     string
	clientsecret string
	code         string
	granttype    string
	endpoint     string
}

type MiniApp interface {
	GetUserInfo() (MiniAppUserInfo, error)
}

func New(input MiniAppInput) MiniApp {
	return &mini{
		clientid:     input.ClientId,
		clientsecret: input.ClientSecret,
		granttype:    input.GrantType,
		redirecturi:  input.RedirectUri,
		code:         input.Code,
		endpoint:     "https://z-wallet.monpay.mn/v2",
	}
}

func (m *mini) GetUserInfo() (MiniAppUserInfo, error) {
	var userInfo MiniAppUserInfo
	err := m.miniApphttpRequest(nil, MiniAppGetUserInfo, "", &userInfo)
	return userInfo, err
}

func (m *mini) CreateInvoice(amount float64, clientServiceUrl, receiver, description string, invoiceType InvoiceType) (response CreateInvoiceRequest, err error) {
	request := CreateInvoiceRequest{
		Amount:           amount,
		RedirectUri:      m.redirecturi,
		ClientServiceUrl: clientServiceUrl,
		Receiver:         receiver,
		InvoiceType:      string(invoiceType),
		Description:      description,
	}

	err = m.miniApphttpRequest(&request, MiniAppCreateInvoice, "", &response)
	return
}

func (m *mini) CheckInvoice()
