package models

type Invoice struct {
	ID               int    `xml:"id"`
	Status           string `xml:"status"`
	InvoiceID        int    `xml:"invois_id"`
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