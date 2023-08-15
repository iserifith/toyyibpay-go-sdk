package toyyibpay

type BillTransaction struct {
	BillName                  string `json:"billName"`
	BillDescription           string `json:"billDescription"`
	BillTo                    string `json:"billTo"`
	BillEmail                 string `json:"billEmail"`
	BillPhone                 string `json:"billPhone"`
	BillStatus                string `json:"billStatus"`
	BillPaymentStatus         string `json:"billpaymentStatus"`
	BillPaymentChannel        string `json:"billpaymentChannel"`
	BillPaymentAmount         string `json:"billpaymentAmount"`
	BillPaymentInvoiceNo      string `json:"billpaymentInvoiceNo"`
	BillSplitPayment          string `json:"billSplitPayment"`
	BillPaymentSettlement     string `json:"billpaymentSettlement"`
	BillPaymentSettlementDate string `json:"billpaymentSettlementDate"`
	BillPaymentDate           string `json:"billPaymentDate"`
	BillExternalReferenceNo   string `json:"billExternalReferenceNo"`
}
