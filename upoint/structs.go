package upoint

type (
	UpointCheckInfoRequest struct {
		CardNumber string `json:"card_number"` // Лояалти системд бүртгэгдсэн картын дугаар
		Mobile     string `json:"mobile"`      // Хэрэглэгчийн утасны дугаар ( хоосон илгээх боломжтой)
		PinCode    string `json:"pin_code"`    // Хэрэглэгчийн пин код
	}

	UpointCheckInfoResponse struct {
		CardStatus int     `json:"card_status"` // Төлөв 1: Идэвхтэй карт (Утасны дугаар оруулахгүй) 2: Шинэ карт (Утасны дугаар оруулах) 3: Хүчингүй карт 4: Буруу карт
		CreatedAt  string  `json:"created_at"`  // Хэрэглэгчийн бүртгэгдсэн огноо
		UAID       string  `json:"ua_id"`       // Хэрэглэгчийн дахин давтагдашгүй дугаар
		CardNumber string  `json:"card_number"` // Хэрэглэгчийн картын дугаар
		Mobile     string  `json:"mobile"`      // Хэрэглэгчийн утасны дугаар
		Balance    float64 `json:"balance"`     // Хэрэглэгчийн онооны үлдэгдэл
		Result     int     `json:"result"`      // Алдааны мэдээллийн код
		Message    string  `json:"message"`     // Алдааны мэдээллийн тайлбар
	}

	UpointQrCheckInfoResponse struct {
		UAID       string  `json:"ua_id"`
		Mobile     string  `json:"modile"`
		Message    string  `json:"message"`
		Balance    float64 `json:"balance"`
		CardStatus int     `json:"card_status"`
		CreatedAt  string  `json:"created_at"`
		CardNumber string  `json:"card_number"`
		Result     int     `json:"result"`
	}

	UpointBankRequest struct {
		BankCode      string  `json:"bank_code"`       // Банкны код
		NonCashAmount float64 `json:"non_cash_amount"` // Бэлэн бусаар төлсөн дүн
	}
	UpointBankResponse struct {
		BankCode      string `json:"bank_code"`       // Банкны код
		Point         string `json:"point"`           // Тухайн гүйлгээнд өгсөн оноо
		NonCashAmount string `json:"non_cash_amount"` // Бэлэн бусаар төлсөн дүн
	}
	UpointManufacturerRequest struct {
		ManufacturerCode   string  `json:"manufacturer_code"`   // Үйлдвэрлэгчийн код
		ManufacturerAmount float64 `json:"manufacturer_amount"` // Үйлдвэрлэгчийн барааны нийт дүн
	}
	UpointManufacturerResponse struct {
		ManufacturerCode   string `json:"manufacturer_code"`   // Үйлдвэрлэгчийн код
		ManufacturerAmount string `json:"manufacturer_amount"` // Үйлдвэрлэгчийн барааны нийт дүн
		Point              string `json:"point"`               // Үйлдвэрлэгчээс үүссэн нийт оноо
	}
	UpointItemRequest struct {
		Code       string  `json:"code"`        // Зураасан код эсвэл байгууллагын дотоод код
		Name       string  `json:"name"`        // Барааны нэр
		Unit       string  `json:"unit"`        // Хэмжих нэгж
		Quantity   float64 `json:"quantity"`    // Тоо, хэмжээ
		Price      float64 `json:"price"`       // Нэгж үнэ
		TotalPrice float64 `json:"total_price"` // Нийт үнэ
	}
	UpointItemResponse struct {
		Barcode   string  `json:"barcode"`    // Зураасан код эсвэл байгууллагын дотоод код
		Point     string  `json:"point"`      // Тухайн бүтээгдэхүүн дээр өгсөн оноо
		Quantity  float64 `json:"qty"`        // Тоо, хэмжээ
		UnitPoint string  `json:"unit_point"` // Хэмжих нэгж
	}
	UpointTransactionRequest struct {
		CardNumber   string                      `json:"card_number"`  // U-Point картын дугаар ***
		InterNumber  string                      `json:"inter_number"` // Юнителийн гэрээний дугаар
		Mobile       string                      `json:"mobile"`       // Хэрэглэгчийн утасны дугаар (Хэрэглэгчид шинээр карт  олгосон үед утгатай байна, бусад үед null байна)***
		Date         string                      `json:"date"`         // Биллийн огноо (YYYY/MM/DD HH:mm:ss) ***
		BillNumber   string                      `json:"bill_number"`  // Биллийн дугаар ***
		SpendAmount  float64                     `json:"spend_amount"` // Зарцуулах онооны дүн ( 0-р илгээх боломжтой) ***
		BonusAmount  float64                     `json:"bonus_amount"` // Оноо үүсгэх дүн ( 0-р илгээх боломжтой) ***
		TotalAmount  float64                     `json:"total_amount"` // Нийт дүн ( 0-р илгээх боломжтой) ***
		CashAmount   float64                     `json:"cash_amount"`  // Бэлнээр төлсөн дүн ( 0-р илгээх боломжтой)
		TerminalID   string                      `json:"terminal_id"`  // ПОС төхөөрөмжийн давтагдахгүй дугаар
		Bank         []UpointBankRequest         `json:"bank"`         // Бэлэн бусаар төлсөн мэдээлэл (хоосон илгээх боломжтой) ***
		Manufacturer []UpointManufacturerRequest `json:"manufacturer"` // Үйлдвэрлэгч болон Импортлогчийн урамшуулал (хоосон илгээх боломжтой) ***
		Items        []UpointItemRequest         `json:"items"`        // Урамшуулсан бүтээгдэхүүний мэдээллийн жагсаалт (хоосон  илгээх боломжтой) ***
	}
	UpointTransactionResponse struct {
		ReceiptID              string                       `json:"receipt_id"`               // U-Point гүйлгээний ID
		Date                   string                       `json:"date"`                     // Биллийн огноо (YYYY/MM/DD HH:mm:ss)
		CardNumber             string                       `json:"card_number"`              // U-Point картын дугаар
		PointBalance           float64                      `json:"point_balance"`            // Хэрэглэгчийн онооны үлдэгдэл
		TotalPoint             float64                      `json:"total_point"`              // Гүйлгээнээс үүссэн оноо
		MerchantPoint          float64                      `json:"merchant_point"`           // Үнийн дүнгээс бодогдож өгсөн оноо
		ManufacturerItemsPoint float64                      `json:"manufacturer_items_point"` // Үйлдвэрлэгч болон импортлогчоос өгсөн оноо
		SpendPoint             float64                      `json:"spend_point"`              // Зарцуулагдсан онооны дүн
		BillNumber             string                       `json:"bill_number"`              // Биллийн дугаар
		Bank                   []UpointBankResponse         `json:"bank"`                     // Бэлэн бус гүйлгээ хийсэн тохиолдолд банкуудаас өгөх оноо
		Items                  []UpointItemResponse         `json:"items"`                    // Борлуулсан бүтээгдэхүүний мэдээллийн жагсаалт
		ManufacturerPoint      float64                      `json:"manufacturer_point"`       // Үйлдвэрлэгч болон импортлогчоос урамшуулалт оноо
		Manufacturer           []UpointManufacturerResponse `json:"manufacturer"`             // Үйлдвэрлэгч
		Result                 float64                      `json:"result"`                   // Алдааны мэдээллийн код
		Message                string                       `json:"message"`                  // Алдааны мэдээллийн тайлбар
		BankPoint              float64                      `json:"bank_point"`               // Банкнаас үүссэн оноо
	}
	UpointReturnTransactionResponse struct {
		InvoiceUUID        string                        `json:"invoice_uuid"`        // Нэхэмжлэхийн uuid
		ReceiptID          string                        `json:"receipt_id"`          // U-Point гүйлгээний ID
		ReturnReceiptID    float64                       `json:"return_receipt_id"`   // Буцаалтын талоны ID
		ManufacturerAmount float64                       `json:"manufacturer_amount"` // Үйлдвэрлэгчийн барааны нийт дүн
		RefundSpendPoint   float64                       `json:"refund_spend_point"`  // Өмнөх гүйлгээнд зарцуулагдсан онооноос буцаах дүн
		RefundBonusAmount  float64                       `json:"refund_bonus_point"`  // Өмнөх гүйлгээнээс нэмэгдсэн онооноос буцаах дүн
		Result             int                           `json:"result"`              // Алдааны мэдээллийн код
		Message            string                        `json:"message"`             // Алдааны мэдээллийн тайлбар
		BillNumber         string                        `json:"bill_number"`         // Биллийн дугаар
		PointBalance       float64                       `json:"point_balance"`       // Хэрэглэгчийн U-point үлдэгдэл
		ItemAmount         float64                       `json:"item_amount"`         // Буцаалт хийсэн талоны бараанаас үүссэн оноо
		BankAmount         float64                       `json:"bank_amount"`         // Буцаалт хийсэн талоны банкнаас үүссэн оноо
		Bank               []*UpointBankResponse         `json:"bank"`                // Бэлэн бус гүйлгээ хийсэн тохиолдолд банкны мэдээлэл
		Manufacturer       []*UpointManufacturerResponse `json:"manufacturer"`        // Үйлдвэрлэгч болон Импортлогчийн урамшуулал
		Items              []*UpointItemResponse         `json:"items"`               // Борлуулсан бүтээгдэхүүний мэдээллийн жагсаалт
	}
	UpointTransactionQrRequest struct {
		QrString     string                      `json:"qr_string"`    // Qr текст
		Date         string                      `json:"date"`         // Биллийн огноо (YYYY/MM/DD HH:mm:ss) ***
		BillNumber   string                      `json:"bill_number"`  // Биллийн дугаар ***
		SpendAmount  float64                     `json:"spend_amount"` // Зарцуулах онооны дүн ( 0-р илгээх боломжтой) ***
		BonusAmount  float64                     `json:"bonus_amount"` // Оноо үүсгэх дүн ( 0-р илгээх боломжтой) ***
		TotalAmount  float64                     `json:"total_amount"` // Нийт дүн ( 0-р илгээх боломжтой) ***
		CashAmount   float64                     `json:"cash_amount"`  // Бэлнээр төлсөн дүн ( 0-р илгээх боломжтой)
		TerminalID   string                      `json:"terminal_id"`  // ПОС төхөөрөмжийн давтагдахгүй дугаар
		BonusPoint   float64                     `json:"bonus_point"`
		Percent      float64                     `json:"percent"`
		Bank         []UpointBankRequest         `json:"bank"`         // Бэлэн бусаар төлсөн мэдээлэл (хоосон илгээх боломжтой) ***
		Manufacturer []UpointManufacturerRequest `json:"manufacturer"` // Үйлдвэрлэгч болон Импортлогчийн урамшуулал (хоосон илгээх боломжтой) ***
		Items        []UpointItemRequest         `json:"items"`        // Урамшуулсан бүтээгдэхүүний мэдээллийн жагсаалт (хоосон  илгээх боломжтой) ***
	}
	UpointTransactionQrResponse struct {
		ReceiptID              string                       `json:"receipt_id"`               // U-Point гүйлгээний ID
		Date                   string                       `json:"date"`                     // Биллийн огноо (YYYY/MM/DD HH:mm:ss)
		CardNumber             string                       `json:"card_number"`              // U-Point картын дугаар
		PointBalance           float64                      `json:"point_balance"`            // Хэрэглэгчийн онооны үлдэгдэл
		TotalPoint             float64                      `json:"total_point"`              // Гүйлгээнээс үүссэн оноо
		MerchantPoint          float64                      `json:"merchant_point"`           // Үнийн дүнгээс бодогдож өгсөн оноо
		ManufacturerItemsPoint float64                      `json:"manufacturer_items_point"` // Үйлдвэрлэгч болон импортлогчоос өгсөн оноо
		SpendPoint             float64                      `json:"spend_point"`              // Зарцуулагдсан онооны дүн
		BillNumber             string                       `json:"bill_number"`              // Биллийн дугаар
		Bank                   []UpointBankResponse         `json:"bank"`                     // Бэлэн бус гүйлгээ хийсэн тохиолдолд банкуудаас өгөх оноо
		Items                  []UpointItemResponse         `json:"items"`                    // Борлуулсан бүтээгдэхүүний мэдээллийн жагсаалт
		ManufacturerPoint      float64                      `json:"manufacturer_point"`       // Үйлдвэрлэгч болон импортлогчоос урамшуулалт оноо
		Manufacturer           []UpointManufacturerResponse `json:"manufacturer"`             // Үйлдвэрлэгч
		Result                 int                          `json:"result"`                   // Алдааны мэдээллийн код
		Message                string                       `json:"message"`                  // Алдааны мэдээллийн тайлбар
		BankPoint              int                          `json:"bank_point"`               // Банкнаас үүссэн оноо
	}
	UpointReturnTransactionRequest struct {
		ReceiptID         string                      `json:"receipt_id"`          // U-Point гүйлгээний ID
		RefundSpendAmount float64                     `json:"refund_spend_amount"` // Өмнөх гүйлгээнд зарцуулагдсан онооноос буцаах дүн
		RefundBonusAmount float64                     `json:"refund_bonus_amount"` // Өмнөх гүйлгээнээс нэмэгдсэн онооноос буцаах дүн
		RefundCashAmount  float64                     `json:"refund_cash_amount"`  // Гүйлгээнд үүссэн онооноос буцаах боломжгүй болсон оноог  мөнгөн дүнгээр авсан дүн
		TerminalID        string                      `json:"terminal_id"`         // ПОС төхөөрөмжийн давтагдахгүй дугаар
		Bank              []UpointBankRequest         `json:"bank"`                // Бэлэн бус гүйлгээ хийсэн тохиолдолд банкны мэдээлэл
		Manufacturer      []UpointManufacturerRequest `json:"manufacturer"`        // Үйлдвэрлэгч болон Импортлогчийн урамшуулал
		Items             []UpointItemRequest         `json:"items"`               // Борлуулсан бүтээгдэхүүний мэдээллийн жагсаалт
	}
	UpointCheckTransactionRequest struct {
		BillNumber string `json:"bill_number"` // Картаар үйлчлүүлэгч байгууллагын гүйлгээний дугаар
		TerminalID string `json:"terminal_id"` // ПОС төхөөрөмжийн давтагдахгүй дугаар
	}
	UpointCheckTransactionResponse struct {
		ReceiptID              string                       `json:"receipt_id"`               // U-Point гүйлгээний ID
		Date                   string                       `json:"date"`                     // Биллийн огноо (YYYY/MM/DD HH:mm:ss)
		CardNumber             string                       `json:"card_number"`              // U-Point картын дугаар
		PointBalance           int                          `json:"point_balance"`            // Хэрэглэгчийн онооны үлдэгдэл
		TotalPoint             int                          `json:"total_point"`              // Гүйлгээнээс үүссэн оноо
		MerchantPoint          int                          `json:"merchant_point"`           // Үнийн дүнгээс бодогдож өгсөн оноо
		ManufacturerItemsPoint float64                      `json:"manufacturer_items_point"` // Үйлдвэрлэгч болон импортлогчоос өгсөн оноо
		SpendPoint             float64                      `json:"spend_point"`              // Зарцуулагдсан онооны дүн
		BillNumber             string                       `json:"bill_number"`              // Биллийн дугаар
		Bank                   []UpointBankResponse         `json:"bank"`                     // Бэлэн бус гүйлгээ хийсэн тохиолдолд банкуудаас өгөх оноо
		Items                  []UpointItemResponse         `json:"items"`                    // Борлуулсан бүтээгдэхүүний мэдээллийн жагсаалт
		ManufacturerPoint      float64                      `json:"manufacturer_point"`       // Үйлдвэрлэгч болон импортлогчоос урамшуулалт оноо
		Manufacturer           []UpointManufacturerResponse `json:"manufacturer"`             // Үйлдвэрлэгч
		Result                 int                          `json:"result"`                   // Алдааны мэдээллийн код
		Message                string                       `json:"message"`                  // Алдааны мэдээллийн тайлбар
		BankPoint              int                          `json:"bank_point"`               // Банкнаас үүссэн оноо
	}
	UpointProductRequest struct {
		Username string `json:"username"` // Үйлдвэрлэгч
		Passowrd string `json:"password"` // Нууц үг
	}
	UpointProductResponse struct {
		Barcode string `json:"barcode"` // Зураасан код эсвэл байгууллагын дотоод код
	}
	UpointCancelTransactionRequest struct {
		BillNumber string  `json:"bill_number"` // Биллийн дугаар
		CashAmount float64 `json:"cash_amount"` // Бэлнээр төлсөн дүн ( 0-р илгээх боломжтой)
	}
	UpointCancelTransactionResponse struct {
		ReceiptID         string                        `json:"receipt_id"`          // U-Point гүйлгээний ID
		ReturnReceiptID   float64                       `json:"return_receipt_id"`   // Буцаалтын талоны ID
		Bank              []*UpointBankResponse         `json:"bank"`                // Бэлэн бус гүйлгээ хийсэн тохиолдолд банкны мэдээлэл
		Manufacturer      []*UpointManufacturerResponse `json:"manufacturer"`        // Үйлдвэрлэгч болон Импортлогчийн урамшуулал
		Items             []*UpointItemResponse         `json:"items"`               // Борлуулсан бүтээгдэхүүний мэдээллийн жагсаалт
		ManufacturerPoint float64                       `json:"manufacturer_point"`  // Үйлдвэрлэгчийн барааны нийт дүн
		RefundSpendPoint  float64                       `json:"refund_spend_point"`  // Өмнөх гүйлгээнд зарцуулагдсан онооноос буцаах дүн
		RefundBonusAmount float64                       `json:"refund_bonus_amount"` // Өмнөх гүйлгээнээс нэмэгдсэн онооноос буцаах дүн
		Result            int                           `json:"result"`              // Алдааны мэдээллийн код
		Message           string                        `json:"message"`             // Алдааны мэдээллийн тайлбар
		BillNumber        string                        `json:"bill_number"`         // Биллийн дугаар
		PointBalance      float64                       `json:"point_balance"`       // Хэрэглэгчийн U-point үлдэгдэл
		ItemAmount        float64                       `json:"item_amount"`         // Буцаалт хийсэн талоны бараанаас үүссэн оноо
		BankAmount        float64                       `json:"bank_amount"`         // Буцаалт хийсэн талоны банкнаас үүссэн оноо
	}
	UpointQrResponse struct {
		Code     string `json:"code"`
		IsNew    string `json:"is_new"`
		QrString string `json:"qr_string"`
	}
	UpointQrCheckResponse struct {
		Status     string `json:"status"`
		CardNumber string `json:"card_number"`
	}
)
