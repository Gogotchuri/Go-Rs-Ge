package main

import (
	"fmt"
	"github.com/gogotchuri/go-rs-ge/models"
	"github.com/gogotchuri/go-rs-ge/rs"
	"time"
)

func main() {
	rs := &rs.Client{
		ServiceUser:     "beka:405217420",
		ServicePassword: "RT1234",
		UserID:          0,
	}
	_, err := rs.GetUniqueID()
	if err != nil {
		fmt.Println(err.Error())
		return
	}


	invoice, err := rs.GetInvoice(162614664)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(invoice)

	customer, err := rs.GetTINFromUniqueID(invoice.BuyerUID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(customer)

	customers, err := rs.GetAllCustomers(&models.InvoiceSearchFilters{
			StartDate:          time.Now().AddDate(-3, 0, 0),
			EndDate:            time.Now().AddDate(3, 0, 0),
			OperationStartDate: time.Now().AddDate(-3, 0, 0),
			OperationEndDate:   time.Now().AddDate(3, 0, 0),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(customers)
	for _, cust := range customers {
		fmt.Println(cust.Name, cust.TIN)
	}
	//invoiceItems, err := rs.GetInvoiceItems(162614664)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(invoiceItems)
	//
	//
	//invoices, err := rs.GetSellerInvoices(&models.InvoiceSearchFilters{
	//	StartDate:          time.Now().AddDate(-3, 0, 0),
	//	EndDate:            time.Now().AddDate(3, 0, 0),
	//	OperationStartDate: time.Now().AddDate(-3, 0, 0),
	//	OperationEndDate:   time.Now().AddDate(3, 0, 0),
	//	InvoiceNo:          "",
	//	CompanyTIN:         "",
	//	Desc:               "",
	//})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//for _, inv := range invoices {
	//	fmt.Println(inv)
	//}
}
