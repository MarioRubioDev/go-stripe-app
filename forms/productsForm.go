package forms

import (
	"bufio"
	"fmt"
	"go-stripe-app/stripemethods/products"
	"os"
)

func FormNewProduct() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Nombre : ")
	name, _ := reader.ReadString('\n')

	fmt.Printf("Descripcion : ")
	description, _ := reader.ReadString('\n')

	fmt.Printf("Precio : ")
	var priceValue float32
	fmt.Scanln(&priceValue)
	priceValue = priceValue * 100

	var newProduct products.StripeProduct = products.StripeProduct{
		Name:        name,
		Description: description,
		PriceValue:  int64(priceValue),
	}

	newProductInfo, newPriceInfo := products.AddNewStripeProduct(newProduct)
	println("Se ha creado un nuevo artículo con ID : ", newProductInfo)
	println("Se ha creado un precio con ID : ", newPriceInfo)
}
