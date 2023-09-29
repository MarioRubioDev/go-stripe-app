package products

import (
	"fmt"
	"go-stripe-app/stripemethods"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"github.com/stripe/stripe-go/v75/paymentintent"
	"github.com/stripe/stripe-go/v75/paymentlink"
	"github.com/stripe/stripe-go/v75/price"
	"github.com/stripe/stripe-go/v75/product"
)

var stripeKey string = stripemethods.DotEnvVariable("STRIPE_KEY")

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

func GetAllProducts() *product.Iter {
	stripe.Key = stripeKey
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := product.List(params)
	return i
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
