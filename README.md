# toyyibPay SDK for Go

This SDK provides a Go interface for the toyyibPay API, allowing developers to easily integrate and communicate with toyyibPay services.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Initialization](#initialization)
  - [Create Category](#create-category)
  - [Create Bill](#create-bill)
  - [Get Bill Transactions](#get-bill-transactions)
- [Examples](#examples)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install the toyyibPay SDK, use `go get`:

```bash
go get github.com/iserifith/toyyibpay-sdk
```

Ensure you have Go modules enabled.

## Usage

### Initialization

To start using the SDK, you need to initialize a new client:

```go
import "https://github.com/iserifith/toyyibpay-go-sdk/pkg/toyyibpay"

client := toyyibpay.New("YOUR_SECRET_KEY", true) // true for sandbox mode, false for production
```

### Create Category

```go
category, err := client.CreateCategory("Category Name", "Category Description")
if err != nil {
log.Fatalf("Error creating category: %v", err)
}
fmt.Println("Created category with code:", category.CategoryCode)
```

### Create Bill

```go
bill, err := client.CreateBill("Bill Name", "Bill Description", "CATEGORY_CODE", 1000) // 1000 is the amount
if err != nil {
log.Fatalf("Error creating bill: %v", err)
}
fmt.Println("Created bill with code:", bill.BillCode)
```

### Get Bill Transactions

```go
transactions, err := client.GetBillTransactions("BILL*CODE", "PAYMENT_STATUS")
if err != nil {
log.Fatalf("Error fetching transactions: %v", err)
}
for *, transaction := range transactions {
fmt.Println("Transaction:", transaction.BillName, transaction.BillPaymentAmount)
}
```

## Examples

For more detailed examples on how to use the SDK, check the `cmd/example` directory.

## Testing

To run tests:

```bash
go test ./...
```

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Write tests for your changes.
4. Implement your changes.
5. Ensure all tests pass.
6. Submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
