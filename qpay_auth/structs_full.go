package qpay_auth

import "time"

type (
	QPayInvoiceRequest struct {
		InvoiceCode         string              `json:"invoice_code"`          // qpay-ээс өгсөн нэхэмжлэхийн код
		SenderInvoiceNo     string              `json:"sender_invoice_no"`     // Байгууллагаас үүсгэх давтагдашгүй нэхэмжлэлийн дугаар
		SenderBranchNo      string              `json:"sender_branch_no"`      // 	Байгууллагын салбарын код
		SenderBranchData    SenderBranchData    `json:"sender_branch_data"`    // Байгууллагын салбарын мэдээлэл
		SenderStaffCode     string              `json:"sender_staff_code"`     // Байгууллагын ажилтаны давтагдашгүй код
		SenderStaffData     SenderStaffData     `json:"sender_staff_data"`     // Байгууллагын ажилтаны мэдээлэл
		SenderTerminalCode  string              `json:"sender_terminal_code"`  // Байгууллага өөрсдийн ашиглаж буй терминалаа давхцалгүй дугаарласан код
		SenderTerminalData  SenderTerminalData  `json:"sender_terminal_data"`  // Байгууллагын терминалын мэдээлэл
		InvoiceReceiverCode string              `json:"invoice_receiver_code"` // Байгууллагын нэхэмжлэхийг хүлээн авч буй харилцагчийн дахин давтагдашгүй дугаар
		InvoiceRecieverData InvoiceRecieverData `json:"invoice_reciever_data"` // Нэхэмжлэл хүлээн авагчийн мэдээлэл
		InvoiceDescription  string              `json:"invoice_description"`   // Нэхэмжлэлийн утга
		InvoiceDueData      time.Time           `json:"invoice_due_date"`      // Нэхэмжлэлийн хүчингүй болох огноо
		EnableExpiry        bool                `json:"enable_expiry"`         // Хугацаа хэтэрсэн ч төлж болох эсэх
		ExpiryDate          time.Time           `json:"expiry_date"`           // Нэхэмжлэхийн дуусах хугацаа
		CalculateVat        bool                `json:"calculate_vat"`         // Нөат-ын тооцоолол
		TaxCustomerCode     bool                `json:"tax_customer_code"`     // ИБаримт үүсгүүлэх байгууллагын мэдээлэл - Байгууллага бол байгууллагын регистерийн дугаар - Хэрэглэгчийн регистерийн дугаар
		LineTaxCode         string              `json:"line_tax_code"`         // БТҮК код - Барааны Lines хоосон үед ашиглана http://web.nso.mn/meta_sys1/files/angilal/Buteegdexuunii%20angilal.pdf
		AllowPartial        bool                `json:"allow_partial"`         // Хувааж төлж болох эсэх
		MinimumAmount       int64               `json:"minimum_amount"`        // Төлөх хамгийн бага дүн
		AllowExceed         bool                `json:"allow_exceed"`          // Илүү төлж болох
		MaximumAmount       int64               `json:"maximum_amount"`        // Төлөх хамгийн их дүн
		Amount              int64               `json:"amount"`                // Мөнгөн дүн
		CallbackUrl         string              `json:"callback_url"`          // Төлбөр төлсөгдсөн эсэх талаар мэдэгдэл авах URL
		Note                string              `json:"note"`                  // Тэмдэглэл
		Lines               []Line              `json:"lines"`                 // Мөрүүд
		Transactions        []Transaction       `json:"transactions"`          // Гүйлгээ
	}
	SenderBranchData struct {
		Register string  `json:"register"` // Салбарын регистр
		Name     string  `json:"name"`     // Салбарын нэр
		Email    string  `json:"email"`    // Салбарын И-майл
		Phone    string  `json:"phone"`    // Салбарын утас
		Address  Address `json:"address"`  // Хаяг
	}
	Address struct {
		City      string `json:"city"`
		District  string `json:"district"`
		Street    string `json:"street"`
		Building  string `json:"building"`
		Address   string `json:"address"`
		Zipcode   string `json:"zipcode"`
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	}
	SenderStaffData struct {
	}
	SenderTerminalData struct {
		Name string `json:"name"`
	}
	InvoiceRecieverData struct {
		Register string  `json:"register"` // 	Нэхэмжлэгч хүлээж авагчийн регистр
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Phone    string  `json:"phone"`
		Address  Address `json:"address"`
	}
	Line struct {
		SenderProductCode string      `json:"sender_product_code"`
		TaxProductCode    string      `json:"tax_product_code"`
		LineDescription   string      `json:"line_description"`
		LineQuantity      int64       `json:"line_quantity"`
		LineUnitPrice     int64       `json:"line_unit_price"`
		Note              string      `json:"note"`
		Discounts         []Discount  `json:"disctounts"`
		Surcharches       []Surcharge `json:"surcharges"`
		Taxes             []Tax       `json:"taxes"`
	}
	Discount struct {
		DicountCode string `json:"discount_code"`
		Description string `json:"description"`
		Amount      int64  `json:"amount"`
		Note        string `json:"note"`
	}
	Surcharge struct {
		SurchargeCode string `json:"surcharge_code"`
		Description   string `json:"description"`
		Amount        int64  `json:"amount"`
		Note          string `json:"note"`
	}
	Tax struct {
		TaxCode     string `json:"tax_code"`
		Description string `json:"description"`
		Amount      int64  `json:"amount"`
		CityTax     int64  `json:"city_tax"`
		Note        string `json:"note"`
	}
	Transaction struct {
		Description string    `json:"description"`
		Amount      int64     `json:"amount"`
		Accounts    []Account `json:"accounts"`
	}
	Account struct {
		AccountBankCode string `json:"account_bank_code"`
		AccountNumber   string `json:"account_number"`
		AccountName     string `json:"account_name"`
		AccountCurrency string `json:"account_currency"`
	}
)
