package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func formGetProductById() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ID del producto: ")
	productId, _ := reader.ReadString('\n')
	return productId
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

func formNewProduct() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Nombre : ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Descripcion : ")
	description, _ := reader.ReadString('\n')

	newArticleId := addNewStripeArticle(name, description)
	println("Se ha creado un nuevo artículo con ID : ", newArticleId)
	mainMenu()
}
func formDeleteProduct() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ID del producto: ")
	productId, _, _ := reader.ReadLine()
	return string(productId)
}
