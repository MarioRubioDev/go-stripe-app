package main

import (
	"bufio"
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/price"
	"github.com/stripe/stripe-go/v75/product"
	"os"
	"strconv"
)

var stripeKey string = "sk_test_51NuCt8AtwzxFIRkS8szHAlTfX3vcwu6gBA13z3uiFG3i6y3DcUrJF7yKh9ez4ZG1KR033dLlYb62bKPElgHgsbKU00zTU41ZOg"

func main() {

	menu()

	//var stripeId string = addNewStripeArticle()
	//setProductPrice(stripeId)
}

func menu() {
	fmt.Println("===================================")
	fmt.Println("== Gestion de articulo en Stripe ==")
	fmt.Println("===================================")
	fmt.Println("1 - Lista de artículos")
	fmt.Println("2 - Añadir nuevo artículo")
	fmt.Println("3 - Listar todos los precios")
	fmt.Println("===================================")
	fmt.Print("Seleccionar opción : ")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		listAllArticles()
	case 2:
		formNewArticle()
	case 3:
		getAllPrices()

	}
}

func formNewPrice() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("ID del Artículo : ")
	stripeId, _ := reader.ReadString('\n')
	fmt.Printf("Import en centimos : ")
	price, _ := reader.ReadString('\n')
	var newPrice int64
	newPrice, _ = strconv.ParseInt(price, 10, 64)
	setProductPrice(stripeId, newPrice)

}

func formNewArticle() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Nombre : ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Descripcion : ")
	description, _ := reader.ReadString('\n')

	newArticleId := addNewStripeArticle(name, description)
	println("Se ha creado un nuevo artículo con ID : ", newArticleId)
	menu()
}

func listAllArticles() {
	stripe.Key = stripeKey

	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := product.List(params)
	for i.Next() {
		p := i.Product()
		fmt.Println("-------------------------------")
		fmt.Println("ID Articulo : ", p.ID)
		fmt.Println("Nonmbre : ", p.Name)
		fmt.Println("Descripción : ", p.Description)
		fmt.Println("ID Precio : ", p.DefaultPrice)
	}
	menu()

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
	menu()
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
