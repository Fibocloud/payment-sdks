package ebarimt3

// ReceiptType defines the type of receipt (bill vs invoice, B2C vs B2B).
type ReceiptType string

// TaxType defines the VAT treatment of a sub-receipt.
type TaxType string

// BarCodeType defines the barcode format for an item.
type BarCodeType string

// PaymentCode defines the payment method.
type PaymentCode string

// PaymentStatus defines the payment result state.
type PaymentStatus string

const (
	// Receipt types
	B2C_RECEIPT = ReceiptType("B2C_RECEIPT") // Business to Consumer receipt
	B2B_RECEIPT = ReceiptType("B2B_RECEIPT") // Business to Business receipt
	B2C_INVOICE = ReceiptType("B2C_INVOICE") // Business to Consumer invoice
	B2B_INVOICE = ReceiptType("B2B_INVOICE") // Business to Business invoice

	// Tax types
	VAT_ABLE = TaxType("VAT_ABLE") // Subject to VAT (НӨАТ тооцох)
	VAT_FREE = TaxType("VAT_FREE") // VAT exempt (НӨАТ-аас чөлөөлөгдөх)
	VAT_ZERO = TaxType("VAT_ZERO") // Zero-rated VAT (НӨАТ 0%)
	NOT_VAT  = TaxType("NOT_VAT")  // No VAT (outside Mongolia)

	// Barcode types
	BARCODE_UNDEFINED = BarCodeType("UNDEFINED")
	BARCODE_GS1       = BarCodeType("GS1")
	BARCODE_ISBN      = BarCodeType("ISBN")

	// Payment codes
	PAYMENT_CASH = PaymentCode("CASH")
	PAYMENT_CARD = PaymentCode("PAYMENT_CARD")

	// Payment statuses
	STATUS_PAID     = PaymentStatus("PAID")
	STATUS_PAY      = PaymentStatus("PAY")
	STATUS_REVERSED = PaymentStatus("REVERSED")
	STATUS_ERROR    = PaymentStatus("ERROR")

	// Receipt response statuses
	RECEIPT_SUCCESS = "SUCCESS"
	RECEIPT_ERROR   = "ERROR"
	RECEIPT_PAYMENT = "PAYMENT"

	// Defaults
	DefaultDistrictCode = "0101" // Ulaanbaatar city center
	DefaultBranchNo     = "001"
)
