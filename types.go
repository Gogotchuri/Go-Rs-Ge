package gorsge

type RsGeClient struct {
	ServiceUser     string
	ServicePassword string
	UserID          int
}

type Operation string
type DeclarationStatus int
type InvoiceStatus int

type Invoice struct {
	Status                 string            `xml:"status"`
	InvoiceID              int               `xml:"invois_id"`
	InvoiceSeries          string            `xml:"f_series"`
	InvoiceNumberF         int               `xml:"f_number"`
	OperationDate          string            `xml:"operation_dt"`
	RegistrationDate       string            `xml:"reg_dt"`
	SellerUID              int               `xml:"seller_un_id"`
	BuyerUID               int               `xml:"buyer_un_id"`
	BuyerSequentialNumber  string            `xml:"seq_num_b"`
	SellerSequentialNumber string            `xml:"seq_num_s"`
	CorrectedInvoiceID     int               `xml:"k_id"`
	CancellerUID           int               `xml:"r_un_id"`
	BuyerServiceUserID     int               `xml:"b_s_user_id"`
	DeclarationStatus      DeclarationStatus `xml:"dec_status"`
}

type CredentialCheck struct {
	Result bool `xml:"chekResult"`
	UserID int  `xml:"user_id"`
	Sua    int  `xml:"sua"`
}
