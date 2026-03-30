package ebarimt3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type client struct {
	endpoint string
}

// EBarimt3 is the interface for the PosAPI 3.0 client.
type EBarimt3 interface {
	// PutReceipt creates a new receipt via POST /rest/receipt.
	PutReceipt(input *CreateReceiptInput) (*CreateReceiptResponse, error)
	// ReturnReceipt cancels a receipt via DELETE /rest/receipt.
	ReturnReceipt(id, date string) error
	// GetInfo returns PosAPI runtime information via GET /rest/info.
	GetInfo() (*InfoResponse, error)
	// SendData manually triggers data sync to the national system via GET /rest/sendData.
	SendData() error
}

// New creates a PosAPI 3.0 client pointed at the given endpoint (e.g. "http://localhost:7080").
func New(endpoint string) EBarimt3 {
	return &client{endpoint: endpoint}
}

// ─── helpers ─────────────────────────────────────────────────────────────────

func f64Ptr(f float64) *float64 { return &f }

func strPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func emptyStrPtr() *string {
	s := ""
	return &s
}

// inputToRequest converts the developer-friendly input into the 3.0 wire format.
func inputToRequest(input *CreateReceiptInput) *CreateReceiptRequest {
	if input.DistrictCode == "" {
		input.DistrictCode = DefaultDistrictCode
	}
	if input.BranchNo == "" {
		input.BranchNo = DefaultBranchNo
	}
	if input.TaxType == "" {
		input.TaxType = VAT_ABLE
	}

	var totalAmount, totalVAT, totalCityTax float64
	var items []Item

	for _, s := range input.Stocks {
		// totalAmount per item = base + vat + cityTax
		itemTotal := s.UnitPrice*s.Qty + s.Vat + s.CityTax
		totalAmount += itemTotal
		totalVAT += s.Vat
		totalCityTax += s.CityTax

		classCode := s.ClassificationCode
		if classCode == "" {
			classCode = input.ClassificationCode
		}

		barcodeType := s.BarCodeType
		if barcodeType == "" {
			if s.BarCode == "" {
				barcodeType = BARCODE_UNDEFINED
			} else {
				barcodeType = BARCODE_GS1
			}
		}

		items = append(items, Item{
			Name:               s.Name,
			BarCode:            strPtrOrNil(s.BarCode),
			BarCodeType:        barcodeType,
			ClassificationCode: classCode,
			TaxProductCode:     strPtrOrNil(s.TaxProductCode),
			MeasureUnit:        s.MeasureUnit,
			Qty:                s.Qty,
			UnitPrice:          s.UnitPrice,
			TotalVAT:           f64Ptr(s.Vat),
			TotalCityTax:       f64Ptr(s.CityTax),
			TotalAmount:        itemTotal,
		})
	}

	// Determine receipt type and buyer identifiers
	receiptType := input.Type
	var customerTin *string
	var consumerNo *string

	if receiptType == "" {
		if input.CustomerTin != "" {
			receiptType = B2B_RECEIPT
		} else {
			receiptType = B2C_RECEIPT
		}
	}

	if receiptType == B2B_RECEIPT || receiptType == B2B_INVOICE {
		customerTin = strPtrOrNil(input.CustomerTin)
	} else {
		consumerNo = strPtrOrNil(input.ConsumerNo)
	}

	subReceipts := []SubReceipt{
		{
			TotalAmount:   totalAmount,
			TotalVAT:      f64Ptr(totalVAT),
			TotalCityTax:  f64Ptr(totalCityTax),
			TaxType:       input.TaxType,
			MerchantTin:   input.MerchantTin,
			CustomerTin:   customerTin,
			BankAccountNo: emptyStrPtr(),
			IBan:          emptyStrPtr(),
			Items:         items,
		},
	}

	payments := []Payment{
		{
			Code:       PAYMENT_CASH,
			Status:     STATUS_PAID,
			PaidAmount: totalAmount,
		},
	}

	return &CreateReceiptRequest{
		BranchNo:     input.BranchNo,
		TotalAmount:  totalAmount,
		TotalVAT:     f64Ptr(totalVAT),
		TotalCityTax: f64Ptr(totalCityTax),
		DistrictCode: input.DistrictCode,
		MerchantTin:  input.MerchantTin,
		PosNo:        input.PosNo,
		CustomerTin:  customerTin,
		ConsumerNo:   consumerNo,
		Type:         receiptType,
		BillIdSuffix: input.BillIDSuffix,
		Receipts:     subReceipts,
		Payments:     payments,
	}
}

// ─── EBarimt3 implementation ─────────────────────────────────────────────────

func (c *client) PutReceipt(input *CreateReceiptInput) (*CreateReceiptResponse, error) {
	if input.MerchantTin == "" {
		return nil, errors.New("MerchantTin is required")
	}
	if input.PosNo == "" {
		return nil, errors.New("PosNo is required")
	}
	if (input.Type == B2B_RECEIPT || input.Type == B2B_INVOICE) && input.CustomerTin == "" {
		return nil, errors.New("CustomerTin is required for B2B receipts")
	}

	body := inputToRequest(input)
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.endpoint+"/rest/receipt", bytes.NewReader(reqBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp CreateReceiptResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	if resp.Status == RECEIPT_ERROR {
		return nil, fmt.Errorf("ebarimt error: %s", resp.Message)
	}

	return &resp, nil
}

func (c *client) ReturnReceipt(id, date string) error {
	body := ReturnReceiptRequest{ID: id, Date: date}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", c.endpoint+"/rest/receipt", bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("return receipt failed with status %d", res.StatusCode)
	}
	return nil
}

func (c *client) GetInfo() (*InfoResponse, error) {
	res, err := http.Get(c.endpoint + "/rest/info")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp InfoResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *client) SendData() error {
	res, err := http.Get(c.endpoint + "/rest/sendData")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
