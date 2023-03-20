package pass

const (
	OrderInqueryStatusPaid      string = "paid"
	OrderInqueryStatusPending   string = "pending"
	OrderInqueryStatusCancelled string = "cancelled"
	OrderInqueryStatusVoided    string = "voided"
)

const (
	WebhookOperationPayment string = "payment"
	WebhookOperationVoid    string = "void"
)
