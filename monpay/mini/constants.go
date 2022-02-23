package mini

type InvoiceType string

const (
	P2P = InvoiceType("P2P")
	P2B = InvoiceType("P2B")
	B2B = InvoiceType("B2B")
)
