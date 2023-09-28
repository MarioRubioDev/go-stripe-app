package products

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/price"
	"github.com/stripe/stripe-go/v75/product"
)

var stripeKey string = "sk_test_51NuCt8AtwzxFIRkS8szHAlTfX3vcwu6gBA13z3uiFG3i6y3DcUrJF7yKh9ez4ZG1KR033dLlYb62bKPElgHgsbKU00zTU41ZOg"

func GetProductById(productId string) *stripe.Product {
	stripe.Key = stripeKey
	p, _ := product.Get(productId, nil)
	return p
}

func GetAllProducts() *product.Iter {
	stripe.Key = stripeKey
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := product.List(params)
	return i
}

func SetProductPrice(id string, newprice int64) {
	stripe.Key = stripeKey
	params := &stripe.PriceParams{
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
		Product:    stripe.String(id),
		UnitAmount: stripe.Int64(newprice),
	}
	p, _ := price.New(params)
	fmt.Println(p)
}

func SetPriceById(priceId string) int64 {
	stripe.Key = stripeKey
	p, _ := price.Get(priceId, nil)
	return p.UnitAmount
}

func GetAllPrices() {
	stripe.Key = stripeKey
	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := price.List(params)
	for i.Next() {
		p := i.Price()
		fmt.Println(p)
	}
}

func AddNewStripeArticle(name string, descrition string) string {
	stripe.Key = stripeKey
	params := &stripe.ProductParams{
		Active:      stripe.Bool(true),
		Description: stripe.String(descrition),
		Name:        stripe.String(name),
		Shippable:   stripe.Bool(false),
		Type:        stripe.String("good"),
	}
	p, _ := product.New(params)
	return p.ID
}

func DeleteProduct(productId string) *stripe.Product {
	stripe.Key = stripeKey
	p, err := product.Del(productId, nil)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return p
}
