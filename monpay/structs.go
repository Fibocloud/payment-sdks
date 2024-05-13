package monpay

import "time"

type (
	MonpayQrInput struct {
		Amount float64
	}
	MonpayQrRequest struct {
		Amount       float64 `json:"amount"`
		DisplayName  string  `json:"string"`
		GenerateUUID bool    `json:"generateUuid"`
		CallbackUrl  string  `json:"callbackUrl"`
	}
	MonpayQrResponse struct {
		Code   int            `json:"code"`
		Info   string         `json:"info"`
		Result MonpayResultQr `json:"result"`
	}
	MonpayResultQr struct {
		Qrcode string `json:"qrcode"`
		UUID   string `json:"uuid"`
	}

	MonpayCheckResponse struct {
		Code   int               `json:"code"`
		Info   string            `json:"info"`
		Result MonpayResultCheck `json:"result"`
	}
	MonpayResultCheck struct {
		UUID          string `json:"uuid"`
		UsedAt        int64  `json:"usedAt"`
		UsedByUd      int64  `json:"usedById"`
		TransactionId string `json:"transactionId"`
		Amount        int64  `json:"amount"`
		CreatedAt     int64  `json:"createdAt"`
		UserPhone     string `json:"userPhone"`
		UserAccountNo string `json:"userAccountNo"`
		UserVatId     string `json:"userVatId"`
		UsedAtUI      string `json:"usedAtUI"`
		CreatedAtUI   string `json:"createdAtUI"`
		AmountUI      string `json:"amountUI"`
	}

	MonpayCallback struct {
		Amount float64 `schema:"amount"`
		UUID   string  `schema:"uuid"`
		Status string  `schema:"status"`
		TnxId  string  `schema:"tnxId"`
	}
)

type (
	DeeplinkCreateRequest struct {
		RedirectUri      string  `json:"redirectUri"`      // Webhook буюу гүйлгээний үр дүн илгээгдэх буцах хаяг
		Amount           float64 `json:"amount"`           // Дүн
		ClientServiceUrl string  `json:"clientServiceUrl"` // Амжилттай гүйлгээний дараа backend-ээс дуудах webhook url.
		// Нэхэмжлэхийн төрлөөс хамаарч утга нь өөр өөр байна.
		//
		// P2B -> Бүтээгдэхүүн борлуулсан салбарын нэр
		//
		// P2P -> Нэхэмжлэгдсэн дүнг хүлээн авах дансны дугаар
		//
		// B2B -> Байгууллагын нэр
		Receiver string `json:"receiver"`
		// P2B ->  Хэрэглэгчээс мерчант
		//
		// P2P -> Хэрэглэгчээс хэрэглэгч
		//
		// B2B -> person to MF & MF to organization
		InvoiceType InvoiceType `json:"invoiceType"`
		Description string      `json:"description"` // Тайлбар
	}
	DeeplinkCreateResponse struct {
		Code    string               `json:"code"`    // Төлөвийн тайлбар код
		IntCode int                  `json:"intCode"` // Төлөвийн код
		Info    string               `json:"info"`    // Төлөвийн мэдээлэл
		Result  DeeplinkCreateResult `json:"result"`
	}

	DeeplinkCreateResult struct {
		ID int `json:"id"` // Нэхэмжлэхийн давтагдашгүй id
		// Нэхэмжлэхийн төрлөөс хамаарч утга нь өөр өөр байна.
		//
		// P2B -> Бүтээгдэхүүн борлуулсан салбарын нэр
		//
		// P2P -> Нэхэмжлэгдсэн дүнг хүлээн авах дансны дугаар
		//
		// B2B -> Байгууллагын нэр
		Receiver    string    `json:"receiver"`
		Amount      float64   `json:"amount"`      // Дүн
		UserID      int       `json:"userId"`      // Төлөгч хэрэглэгчийн id
		MiniAppID   int       `json:"miniAppId"`   // Мини апп id
		CreateDate  time.Time `json:"createDate"`  // Нэхэмжлэх үүссэн огноо
		UpdateDate  time.Time `json:"updateDate"`  // Нэхэмжлэх засагдсан огноо
		Status      string    `json:"status"`      // Нэхэмжлэх статус
		Description string    `json:"description"` // Нэхэмжлэхийн тайлбар
		RedirectUri string    `json:"redirectUri"` // Веб хөтчид нээгдэх буцах url хаяг. Гүйлгээг дууссаныг мэдэгдэнэ
		// P2B ->  Хэрэглэгчээс мерчант
		//
		// P2P -> Хэрэглэгчээс хэрэглэгч
		//
		// B2B -> person to MF & MF to organization
		InvoiceType InvoiceType `json:"invoiceType"`
	}

	DeeplinkCheckResponse struct {
		Code    string              `json:"code"`    // Төлөвийн тайлбар код
		IntCode int                 `json:"intCode"` // Төлөвийн код
		Info    string              `json:"info"`    // Төлөвийн мэдээлэл
		Result  DeeplinkCheckResult `json:"result"`
	}

	DeeplinkCheckResult struct {
		ID int `json:"id"` // Нэхэмжлэхийн давтагдашгүй id
		// Нэхэмжлэхийн төрлөөс хамаарч утга нь өөр өөр байна.
		//
		// P2B -> Бүтээгдэхүүн борлуулсан салбарын нэр
		//
		// P2P -> Нэхэмжлэгдсэн дүнг хүлээн авах дансны дугаар
		//
		// B2B -> Байгууллагын нэр
		Receiver    string    `json:"receiver"`
		Amount      float64   `json:"amount"`      // Дүн
		UserID      int       `json:"userId"`      // Төлөгч хэрэглэгчийн id
		MiniAppID   int       `json:"miniAppId"`   // Мини апп id
		CreateDate  time.Time `json:"createDate"`  // Нэхэмжлэх үүссэн огноо
		UpdateDate  time.Time `json:"updateDate"`  // Нэхэмжлэх засагдсан огноо
		Status      string    `json:"status"`      // Нэхэмжлэх статус
		Description string    `json:"description"` // Нэхэмжлэхийн тайлбар
		StatusInfo  string    `json:"statusInfo"`  // Хэвлэж болохуйц мэдээлэл
		RedirectUri string    `json:"redirectUri"` // Веб хөтчид нээгдэх буцах url хаяг. Гүйлгээг дууссаныг мэдэгдэнэ
		// P2B ->  Хэрэглэгчээс мерчант
		//
		// P2P -> Хэрэглэгчээс хэрэглэгч
		//
		// B2B -> person to MF & MF to organization
		InvoiceType     InvoiceType `json:"invoiceType"`
		BankName        Bank        `json:"bankName"`         // Банк нэр (Only B2B connections)
		BankAccount     string      `json:"bankAccount "`     // Банкны дансны дугаар (Only B2B connections)
		BankAccountName string      `json:"bankAccountName "` // Данс эзэмшигчийн нэр (Only B2B connections)
	}
	DeeplinkCallback struct {
		Amount    float64 `schema:"amount"`    // Нэхэмжилсэн дүн
		InvoiceId string  `schema:"invoiceId"` // Төлсөн нэхэмжлэхийн id
		Status    string  `schema:"status"`    // PAID, FAILED
		TnxId     string  `schema:"tnxId"`     // Гүйлгээ амжилттай болсон бол гүйлгээний дугаар
		Info      string  `schema:"info"`      // Хүнд уншигдахуйц гүйлгээний үр дүн
	}
)
