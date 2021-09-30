package models

import "time"

type SaveInvoiceItemT struct {
	ID           int     `xml:"id"`
	InvoiceID    int     `xml:"invois_id"`
	Name         string  `xml:"goods"`
	Unit         string  `xml:"g_unit"`
	Quantity     float64 `xml:"g_number"`
	TotalAmount  float64 `xml:"full_amount"`
	VATAmount    float64 `xml:"drg_amount"`
	ExciseAmount float64 `xml:"aqcizi_amount"`
	ExciseID     int     `xml:"akciz_id"`
}

type SaveInvoiceT struct {
	InvoiceID          int    `xml:"invois_id"`
	OperationDate      string `xml:"operation_date"`
	SellerUID          int    `xml:"seller_un_id"`
	BuyerUID           int    `xml:"buyer_un_id"`
	BuyerServiceUserID int    `xml:"b_s_user_id"`
}

type InvoiceSearchFilters struct {
	StartDate          time.Time
	EndDate            time.Time
	OperationStartDate time.Time
	OperationEndDate   time.Time

	InvoiceNo  string
	CompanyTIN string
	Desc       string
}

type InvoiceItem struct {
	ID        int    `xml:"ID"`
	InvoiceID int    `xml:"INV_ID"`
	Name      string `xml:"GOODS"`

	Unit     string  `xml:"G_UNIT"`
	Quantity float64 `xml:"G_NUMBER"`

	TotalAmount  float64 `xml:"FULL_AMOUNT"`
	VATAmount    float64 `xml:"DRG_AMOUNT"`
	SVATAmount   float64 `xml:"SDRG_AMOUNT"`
	ExciseAmount float64 `xml:"AQCIZI_AMOUNT"`

	ExciseID  int `xml:"AKCIS_ID"`
	WaybillID int `xml:"WAYBILL_ID"`
}

type Invoice struct {
	ID               int    `xml:"ID"`
	Status           string `xml:"STATUS"`
	InvoiceSeries    string `xml:"F_SERIES"`
	InvoiceNumberF   int    `xml:"F_NUMBER"`
	OperationDate    string `xml:"OPERATION_DT"`
	RegistrationDate string `xml:"REG_DT"`

	Amount float64 `xml:"TANXA"`
	VAT    float64 `xml:"VAT"`

	AgreementDate string `xml:"AGREE_DATE"`
	AgreedByUser  int    `xml:"AGREE_S_USER_ID"`

	BuyerUID              int    `xml:"BUYER_UN_ID"`
	BuyerServiceUserID    int    `xml:"B_S_USER_ID"`
	BuyerSequentialNumber string `xml:"SEQ_NUM_B"`
	BuyerTaxID            int    `xml:"SA_IDENT_NO"`
	BuyerName             string `xml:"ORG_NAME"`

	SellerUID              int    `xml:"SELLER_UN_ID"`
	SellerSequentialNumber string `xml:"SEQ_NUM_S"`
	ServiceUserID          int    `xml:"S_USER_ID"`

	CancellerUID int `xml:"R_UN_ID"`

	DeclarationStatus  DeclarationStatus `xml:"DEC_STATUS"`
	WasRefused         bool              `xml:"WAS_REF"`
	CorrectedInvoiceID int               `xml:"K_ID"`
	CorrectionType     int               `xml:"K_TYPE"`
}

type InvoiceSingle struct {
	ID               int    `xml:"id"`
	Status           string `xml:"status"`
	InvoiceSeries    string `xml:"f_series"`
	InvoiceNumberF   int    `xml:"f_number"`
	OperationDate    string `xml:"operation_dt"`
	RegistrationDate string `xml:"reg_dt"`

	Amount float64 `xml:"tanxa"`
	VAT    float64 `xml:"vat"`

	AgreementDate string `xml:"agree_date"`
	AgreedByUser  int    `xml:"agree_s_user_id"`

	BuyerUID              int    `xml:"buyer_un_id"`
	BuyerServiceUserID    int    `xml:"b_s_user_id"`
	BuyerSequentialNumber string `xml:"seq_num_b"`
	BuyerTaxID            int    `xml:"sa_ident_no"`
	BuyerName             string `xml:"org_name"`

	SellerUID              int    `xml:"seller_un_id"`
	SellerSequentialNumber string `xml:"seq_num_s"`
	ServiceUserID          int    `xml:"s_user_id"`

	CancellerUID int `xml:"r_un_id"`

	DeclarationStatus  DeclarationStatus `xml:"dec_status"`
	WasRefused         bool              `xml:"was_ref"`
	CorrectedInvoiceID int               `xml:"k_id"`
	CorrectionType     int               `xml:"k_type"`
}

func (inv *InvoiceSingle) GetInvoice() *Invoice {
	return &Invoice{
		ID:                     inv.ID,
		Status:                 inv.Status,
		InvoiceSeries:          inv.InvoiceSeries,
		InvoiceNumberF:         inv.InvoiceNumberF,
		OperationDate:          inv.OperationDate,
		RegistrationDate:       inv.RegistrationDate,
		Amount:                 inv.Amount,
		VAT:                    inv.VAT,
		AgreementDate:          inv.AgreementDate,
		AgreedByUser:           inv.AgreedByUser,
		BuyerUID:               inv.BuyerUID,
		BuyerServiceUserID:     inv.BuyerServiceUserID,
		BuyerSequentialNumber:  inv.BuyerSequentialNumber,
		BuyerTaxID:             inv.BuyerTaxID,
		BuyerName:              inv.BuyerName,
		SellerUID:              inv.SellerUID,
		SellerSequentialNumber: inv.SellerSequentialNumber,
		ServiceUserID:          inv.ServiceUserID,
		CancellerUID:           inv.CancellerUID,
		DeclarationStatus:      inv.DeclarationStatus,
		WasRefused:             inv.WasRefused,
		CorrectedInvoiceID:     inv.CorrectedInvoiceID,
		CorrectionType:         inv.CorrectionType,
	}
}
