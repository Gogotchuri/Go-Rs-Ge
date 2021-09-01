package models

import "time"

type InvoiceSearchFilters struct {
	StartDate          time.Time
	EndDate            time.Time
	OperationStartDate time.Time
	OperationEndDate   time.Time

	InvoiceNo  string
	CompanyTIN string
	Desc       string
}

type Invoice struct {
	ID               int    `xml:"ID"`
	Status           string `xml:"STATUS"`
	InvoiceID        int    `xml:"INVOIS_ID"`
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
