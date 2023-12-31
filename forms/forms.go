package forms

import (
	"bufio"
	"fmt"
	"go-stripe-app/stripemethods/productMethods"
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
	productMethods.SetProductPrice(stripeId, newPrice)

}

func FormDeleteProduct() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ID del producto: ")
	productId, _, _ := reader.ReadLine()
	return string(productId)
}
