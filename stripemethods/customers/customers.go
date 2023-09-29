package customers

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/customer"
)

var stripeKey string = "sk_test_51NuCt8AtwzxFIRkS8szHAlTfX3vcwu6gBA13z3uiFG3i6y3DcUrJF7yKh9ez4ZG1KR033dLlYb62bKPElgHgsbKU00zTU41ZOg"

type CustomerApp struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	TaxId    string `json:"taxid"`
	PaymenID []PayProId
	Address  []CustomerAddress
}

type CustomerAddress struct {
	IdAddress   string `json:"idaddress"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Postal_code string `json:"postal_code"`
	State       string `json:"state"`
}

type PayProId struct {
	CustomerID string `json:"customerID"`
	Provider   string `json:"provider"`
}

func AddNewCustomer(NewCustomerData CustomerApp) {
	CustomerStripeId := AddNewStripeCustomer(NewCustomerData)
	NewCustomerData.PaymenID = []PayProId{
		{CustomerID: CustomerStripeId, Provider: "Stripe"},
	}
	AddNewAppCustomer(NewCustomerData)
}
func AddNewStripeCustomer(c CustomerApp) string {
	return ""
}
func AddNewAppCustomer(data CustomerApp) {

}

func GetAllCustomers() {
	stripe.Key = stripeKey
	params := &stripe.CustomerListParams{}
	r := customer.List(params)
	for r.Next() {
		fmt.Println(r.Customer().ID, " ", r.Customer().Email)
	}
}
