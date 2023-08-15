package main

import (
	"fmt"
	"log"

	"github.com/iserifith/toyyibpay-go-sdk/pkg/toyyibpay"
)

func main() {
	client := toyyibpay.New("YOUR_KEY", true) // true for sandbox, false for production
	// Create a category
	categoryResp, err := client.CreateCategory("Test Category", "This is a test category description")
	if err != nil {
		log.Fatalf("Error creating category: %v", err)
	}
	fmt.Println("Created category with code:", categoryResp.CategoryCode)

	// Create a bill
	billReq := toyyibpay.CreateBillRequest{
		CategoryCode:            categoryResp.CategoryCode,
		BillName:                "Test Bill",
		BillDescription:         "This is a test bill description",
		BillAmount:              1000,
		BillPriceSetting:        1,
		BillPayorInfo:           1,
		BillReturnUrl:           "http://example.com/return",
		BillCallbackUrl:         "http://example.com/callback",
		BillExternalReferenceNo: "mockRef",
		BillTo:                  "John Doe",
		BillEmail:               "john@example.com",
		BillPhone:               "1234567890",
	}

	billResp, err := client.CreateBill(billReq)
	if err != nil {
		log.Fatalf("Error creating bill: %v", err)
	}
	fmt.Println("Created bill with code:", billResp.BillCode)

	// Fetch transactions for the bill
	transactions, err := client.GetBillTransactions(billResp.BillCode, "1")
	if err != nil {
		log.Fatalf("Error fetching transactions: %v", err)
	}

	for _, transaction := range transactions {
		fmt.Printf("Transaction: %s", transaction)
	}

	// Fetch category details
	categoryDetails, err := client.GetCategory(categoryResp.CategoryCode)
	if err != nil {
		log.Fatalf("Error fetching category details: %v", err)
	}
	fmt.Printf("Category: %s", categoryDetails)
}
