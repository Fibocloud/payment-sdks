package monpay

type InvoiceType string

const (
	P2P InvoiceType = "P2P" // Хэрэглэгчээс хэрэглэгч
	P2B InvoiceType = "P2B" // Хэрэглэгчээс мерчант
	B2B InvoiceType = "B2B" // person to MF & MF to organization
)

type Bank string

const (
	BankKhan        Bank = "KHAN"
	BankGolomt      Bank = "GOLOMT"
	BankState       Bank = "STATE"
	BankUlaanbaatar Bank = "ULAANBAATAR"
	BankXac         Bank = "XAC"
	BankCapitron    Bank = "CAPITRON"
	BankArig        Bank = "ARIG"
	BankChinggis    Bank = "CHINGGIS"
	BankBogd        Bank = "BOGD"
	BankCredit      Bank = "CREDIT"
	BankHugjil      Bank = "HUGJIL"
	BankTuriinsan   Bank = "TURIIN_SAN"
)
