package views

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/product"
)

func ViewProduct(product *stripe.Product) {
	fmt.Println("ID del producto: " + product.ID)
	fmt.Println("Nombre : " + product.Name)
}

func ViewAllProducts(listProduct *product.Iter) {
	for listProduct.Next() {
		p := listProduct.Product()
		fmt.Println("-------------------------------")
		fmt.Println("ID Articulo : ", p.ID)
		fmt.Println("Nonmbre : ", p.Name)
		fmt.Println("Descripción : ", p.Description)
		fmt.Println("ID Precio : ", p.DefaultPrice)
	}
}

func ViewConfirmDeleted(product *stripe.Product) {
	fmt.Println("------Producto Eliminado -----------")
	fmt.Println("ID Articulo : ", product.ID)
	fmt.Println("Estado eliminación :", product.Deleted)
}
