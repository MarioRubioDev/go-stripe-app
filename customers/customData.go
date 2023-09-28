package customers

import (
	"encoding/json"
	"fmt"
)

// Example data of new customer.
var TestCustomer CustomerApp = CustomerApp{
	ID:    "0fg98fdhgfdp9hpd98hsdph8vds9p8fhv9p8d",
	Name:  "John Doe",
	Email: "john.doe@example.com",
	Phone: "+34-555-5555",
	TaxId: "32998833A",
	PaymenID: []PayProId{
		{CustomerID: "cu_94h3phqf4uqu3hf1", Provider: "Stripe"},
		{CustomerID: "9asyvas98vyas", Provider: "Paypal"},
	},
	Address: []CustomerAddress{
		{
			IdAddress:   "Casa",
			City:        "City A",
			Country:     "Country A",
			Street1:     "Street 1",
			Street2:     "Street 2",
			Postal_code: "12345",
			State:       "State A",
		},
		{
			IdAddress:   "Trabajo",
			City:        "City B",
			Country:     "Country B",
			Street1:     "Street 3",
			Street2:     "Street 4",
			Postal_code: "54321",
			State:       "State B",
		},
	},
}

func PrintCustomerJson() {
	jsonData, err := json.Marshal(TestCustomer)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
