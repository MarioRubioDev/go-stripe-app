package forms

import (
	"bufio"
	"fmt"
	"go-stripe-app/stripemethods/products"
	"os"
	"strconv"
)

func FormGetProductById() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ID del producto: ")
	productId, _, _ := reader.ReadLine()
	return string(productId)
}

func FormNewPrice() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("ID del Artículo : ")
	stripeId, _ := reader.ReadString('\n')
	fmt.Printf("Import en centimos : ")
	price, _ := reader.ReadString('\n')
	var newPrice int64
	newPrice, _ = strconv.ParseInt(price, 10, 64)
	products.SetProductPrice(stripeId, newPrice)

}

func FormNewProduct() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Nombre : ")
	name, _, _ := reader.ReadLine()
	fmt.Printf("Descripcion : ")
	description, _, _ := reader.ReadLine()
	newArticleId := products.AddNewStripeArticle(string(name), string(description))
	println("Se ha creado un nuevo artículo con ID : ", newArticleId)
}
func FormDeleteProduct() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ID del producto: ")
	productId, _, _ := reader.ReadLine()
	return string(productId)
}
