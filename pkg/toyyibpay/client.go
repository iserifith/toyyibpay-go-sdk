package toyyibpay

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseURL        = "https://toyyibpay.com/index.php/api/"
	sandboxBaseURL = "https://dev.toyyibpay.com/index.php/api/"
)

type Client struct {
	SecretKey string
	BaseURL   string
}

func New(secretKey string, sandbox bool) *Client {
	if sandbox {
		return &Client{SecretKey: secretKey, BaseURL: sandboxBaseURL}
	}
	return &Client{SecretKey: secretKey, BaseURL: baseURL}
}

type CreateCategoryResponse struct {
	CategoryCode string `json:"CategoryCode"`
}

func (c *Client) CreateCategory(catName, catDescription string) (*CreateCategoryResponse, error) {
	data := url.Values{}
	data.Set("catname", catName)
	data.Set("catdescription", catDescription)
	data.Set("userSecretKey", c.SecretKey)

	resp, err := http.PostForm(c.BaseURL+"createCategory", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CreateCategoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

type CreateBillResponse struct {
	BillCode string `json:"BillCode"`
}

type CreateBillRequest struct {
	CategoryCode            string
	BillName                string
	BillDescription         string
	BillAmount              int
	BillPriceSetting        int
	BillPayorInfo           int
	BillReturnUrl           string
	BillCallbackUrl         string
	BillExternalReferenceNo string
	BillTo                  string
	BillEmail               string
	BillPhone               string
	BillSplitPayment        *int
	BillSplitPaymentArgs    *string
	BillPaymentChannel      *int
	BillContentEmail        *string
	BillChargeToCustomer    *int
}

func (c *Client) CreateBill(req CreateBillRequest) (*CreateBillResponse, error) {
	data := url.Values{}
	data.Set("userSecretKey", c.SecretKey)
	data.Set("categoryCode", req.CategoryCode)
	data.Set("billName", req.BillName)
	data.Set("billDescription", req.BillDescription)
	data.Set("billAmount", strconv.Itoa(req.BillAmount))
	data.Set("billPriceSetting", strconv.Itoa(req.BillPriceSetting))
	data.Set("billPayorInfo", strconv.Itoa(req.BillPayorInfo))
	data.Set("billReturnUrl", req.BillReturnUrl)
	data.Set("billCallbackUrl", req.BillCallbackUrl)
	data.Set("billExternalReferenceNo", req.BillExternalReferenceNo)
	data.Set("billTo", req.BillTo)
	data.Set("billEmail", req.BillEmail)
	data.Set("billPhone", req.BillPhone)

	if req.BillSplitPayment != nil {
		data.Set("billSplitPayment", strconv.Itoa(*req.BillSplitPayment))
	}

	if req.BillSplitPaymentArgs != nil {
		data.Set("billSplitPaymentArgs", *req.BillSplitPaymentArgs)
	}

	if req.BillPaymentChannel != nil {
		data.Set("billPaymentChannel", strconv.Itoa(*req.BillPaymentChannel))
	}

	if req.BillContentEmail != nil {
		data.Set("billContentEmail", *req.BillContentEmail)
	}

	if req.BillChargeToCustomer != nil {
		data.Set("billChargeToCustomer", strconv.Itoa(*req.BillChargeToCustomer))
	}

	resp, err := http.PostForm(c.BaseURL+"createBill", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CreateBillResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("no response from API")
	}

	return &result[0], nil
}

func (c *Client) GetBillTransactions(billCode, billPaymentStatus string) ([]interface{}, error) {
	data := url.Values{}
	data.Set("billCode", billCode)
	data.Set("billpaymentStatus", billPaymentStatus)

	resp, err := http.PostForm(c.BaseURL+"getBillTransactions", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []interface{}

	return result, nil

}

type GetCategoryResponse struct {
	CategoryName        string `json:"categoryName"`
	CategoryDescription string `json:"categoryDescription"`
	CategoryStatus      string `json:"categoryStatus"`
}

func (c *Client) GetCategory(categoryCode string) ([]interface{}, error) {
	data := url.Values{}
	data.Set("userSecretKey", c.SecretKey)
	data.Set("categoryCode", categoryCode)

	resp, err := http.PostForm(c.BaseURL+"getCategory", data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []interface{}

	return result, nil

}
