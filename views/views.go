package views

import (
	"fmt"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/product"
	"go-stripe-app/stripemethods/productMethods"
)

func ViewProduct(product *stripe.Product) {
	fmt.Println("<------------------------------->")
	fmt.Println("ID del producto: ", product.ID)
	fmt.Println("Nombre : ", product.Name)
	fmt.Println("Precio : ", float64(productMethods.GetProductPriseById(product.ID))/100)
	fmt.Println("<------------------------------->")
}

func ViewAllProducts(listProduct *product.Iter) {
	fmt.Println("<------------------------------->")
	fmt.Printf("ID Producto\t Nombre\t Descripcion\t Precio\n")
	for listProduct.Next() {
		p := listProduct.Product()
		productPrice := float64(productMethods.GetProductPriseById(p.ID) / 100)
		fmt.Printf("%s\t %s\t %s\t %f\n", p.ID, p.Name, p.Description, productPrice)
	}
}

func ViewConfirmDeleted(product *stripe.Product) {
	fmt.Println("------Producto Eliminado -----------")
	fmt.Println("ID Articulo : ", product.ID)
	fmt.Println("Estado eliminación :", product.Deleted)
	fmt.Println("<------------------------------->")
}
