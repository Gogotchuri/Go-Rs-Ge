package requests

const ServiceAddress = "https://www.revenue.mof.ge/ntosservice/ntosservice.asmx?op="

//Operations
const (
	WhatIsMyIP Operation = "what_is_my_ip"

	CheckCredentials   Operation = "chek"
	GetUniqueID        Operation = "get_un_id_from_user_id"
	GetTINFromUniqueID Operation = "get_tin_from_un_id"
	GetUniqueIDFromTIN Operation = "get_un_id_from_tin"

	GetInvoice        Operation = "get_invoice"
	GetSellerInvoices Operation = "get_seller_invoices"
	GetBuyerInvoices  Operation = "get_buyer_invoices"
	GetInvoiceItems   Operation = "get_invoice_desc"
)
