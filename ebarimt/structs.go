package ebarimt

type CreateEbarimtRequest struct {
	Amount        string  `json:"amount"`
	Vat           string  `json:"vat"`
	CashAmount    string  `json:"cashAmount"`
	NonCashAmount string  `json:"nonCashAmount"`
	CityTax       string  `json:"cityTax"`
	CustomerNo    string  `json:"customerNo"`
	BillType      string  `json:"billType"`
	BranchNo      string  `json:"branchNo"`
	DistrictCode  string  `json:"districtCode"`
	Stocks        []Stock `json:"stocks"`
}

type Stock struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	MeasureUnit string `json:"measureUnit"`
	Qty         string `json:"qty"`
	UnitPrice   string `json:"unitPrice"`
	TotalAmount string `json:"totalAmount"`
	CityTax     string `json:"cityTax"`
	Vat         string `json:"vat"`
	BarCode     string `json:"barCode"`
}

type CreateEbarimtResponse struct {
	Amount        string  `json:"amount"`
	Vat           string  `json:"vat"`
	CashAmount    string  `json:"cashAmount"`
	NonCashAmount string  `json:"nonCashAmount"`
	CityTax       string  `json:"cityTax"`
	CustomerNo    string  `json:"customerNo"`
	BillType      string  `json:"billType"`
	BranchNo      string  `json:"branchNo"`
	DistrictCode  string  `json:"districtCode"`
	Stocks        []Stock `json:"stocks"`
	TaxType       string  `json:"taxType"`
	RegisterNo    string  `json:"registerNo"`
	BillId        string  `json:"billId"`
	MacAddress    string  `json:"macAddress"`
	Date          string  `json:"date"`
	Lottery       string  `json:"lottery"`
	InternalCode  string  `json:"internalCode"`
	QrData        string  `json:"qrData"`
	MerchantId    string  `json:"merchantId"`
	Success       bool    `json:"success"`
}

type ReturnBillResponse struct {
	Success bool `json:"success"`
}
