package views

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/product"
	"go-stripe-app/stripemethods/products"
)

func ViewProduct(product *stripe.Product) {
	fmt.Println("<------------------------------->")
	fmt.Println("ID del producto: ", product.ID)
	fmt.Println("Nombre : ", product.Name)
	fmt.Println("Precio : ", float64(products.GetProductPriseById(product.ID))/100)
	fmt.Println("<------------------------------->")
}

func ViewAllProducts(listProduct *product.Iter) {
	for listProduct.Next() {
		p := listProduct.Product()
		fmt.Println("<------------------------------->")
		fmt.Println("ID Articulo : ", p.ID)
		fmt.Println("Nonmbre : ", p.Name)
		fmt.Println("Descripción : ", p.Description)
		fmt.Println("ID Precio : ", float64(products.GetProductPriseById(p.ID))/100)
		fmt.Println("<------------------------------->")
	}
}

func ViewConfirmDeleted(product *stripe.Product) {
	fmt.Println("------Producto Eliminado -----------")
	fmt.Println("ID Articulo : ", product.ID)
	fmt.Println("Estado eliminación :", product.Deleted)
	fmt.Println("<------------------------------->")
}
