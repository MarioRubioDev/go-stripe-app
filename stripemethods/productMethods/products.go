package productMethods

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"github.com/stripe/stripe-go/v75/paymentintent"
	"github.com/stripe/stripe-go/v75/paymentlink"
	"github.com/stripe/stripe-go/v75/price"
	"github.com/stripe/stripe-go/v75/product"
	"go-stripe-app/stripemethods"
)

var stripeKey string = stripemethods.DotEnvVariable("STRIPE_KEY")

type StripeProduct struct {
	Name        string
	Description string
	PriceValue  int64
}

// name: Product name.
// description: Product description.
// priceValue: Product price ( in cents...without decimals).
func AddNewStripeProduct(p StripeProduct) (string, string) {
	stripe.Key = stripeKey
	productParams := &stripe.ProductParams{
		Active:      stripe.Bool(true),
		Name:        stripe.String(p.Name),
		Description: stripe.String(p.Description),
		TaxCode:     stripe.String("txcd_20030000"),
		Type:        stripe.String("service"),
	}
	newProduct, _ := product.New(productParams)

	priceParams := &stripe.PriceParams{
		BillingScheme: stripe.String("per_unit"),
		Currency:      stripe.String(string(stripe.CurrencyEUR)),
		Product:       stripe.String(newProduct.ID),
		UnitAmount:    stripe.Int64(p.PriceValue),
	}
	newPrice, _ := price.New(priceParams)
	return newProduct.ID, newPrice.ID
}

func GetAllProducts() *product.Iter {
	stripe.Key = stripeKey
	params := &stripe.ProductListParams{}
	i := product.List(params)
	return i
}

func GetProductById(productId string) *stripe.Product {
	stripe.Key = stripeKey
	p, _ := product.Get(productId, nil)
	return p
}

func GetProductPriseById(productId string) int64 {
	stripe.Key = stripeKey
	params := &stripe.PriceListParams{
		Product: stripe.String(productId),
	}
	prices := price.List(params)
	productsPrise := prices.PriceList().Data[0].UnitAmount
	return productsPrise
}

func GetPriceById(priceId string) int64 {
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

func GetAllPaymentLinks() {
	stripe.Key = stripeKey
	params := &stripe.PaymentLinkListParams{}
	i := paymentlink.List(params)
	for i.Next() {
		pl := i.PaymentLink().ID
		fmt.Println(pl)
	}
}

func GetAllCheckoutSessions() {
	stripe.Key = stripeKey
	params := &stripe.CheckoutSessionListParams{}
	i := session.List(params)
	for i.Next() {
		s := i.CheckoutSession()
		fmt.Println(s.ID + " " + string(s.PaymentStatus) + " " + string(s.Status))
	}
}

func GetAllPaymentIntent() {
	stripe.Key = stripeKey
	params := &stripe.PaymentIntentListParams{}
	pi := paymentintent.List(params)
	for pi.Next() {
		s := pi.PaymentIntent()
		fmt.Println(s.ID, " ", s.Status, " ", s.AmountReceived, " ", s.ConfirmationMethod)
	}
}

// -----------------------------
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

func DeleteProduct(productId string) *stripe.Product {
	stripe.Key = stripeKey
	p, err := product.Del(productId, nil)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return p
}
