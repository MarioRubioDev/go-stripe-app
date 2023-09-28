package main

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/price"
	"github.com/stripe/stripe-go/v75/product"
)

func getProductById(productId string) *stripe.Product {
	stripe.Key = stripeKey
	p, _ := product.Get(productId, nil)
	return p
}

func getAllProducts() *product.Iter {
	stripe.Key = stripeKey
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := product.List(params)
	return i
}

func setProductPrice(id string, newprice int64) {
	stripe.Key = stripeKey
	params := &stripe.PriceParams{
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
		Product:    stripe.String(id),
		UnitAmount: stripe.Int64(newprice),
	}
	p, _ := price.New(params)
	fmt.Println(p)
}

func getPriceById(priceId string) int64 {
	stripe.Key = stripeKey
	p, _ := price.Get(priceId, nil)
	return p.UnitAmount
}

func getAllPrices() {
	stripe.Key = stripeKey
	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := price.List(params)
	for i.Next() {
		p := i.Price()
		fmt.Println(p)
	}
	mainMenu()
}

func addNewStripeArticle(name string, descrition string) string {
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

func deleteProduct(productId string) *stripe.Product {
	stripe.Key = stripeKey
	p, err := product.Del(productId, nil)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return p
}
