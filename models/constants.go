package models

type InvoiceSatus int

const (
	Deleted              InvoiceSatus = -1
	Created              InvoiceSatus = 0
	Sent                 InvoiceSatus = 1
	Accepted             InvoiceSatus = 2
	Cancelled            InvoiceSatus = 6
	AcceptedAndCancelled InvoiceSatus = 7
)
