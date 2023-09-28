package menus

import (
	"fmt"
	"go-stripe-app/forms"
	"go-stripe-app/products"
	"go-stripe-app/views"
)

func MainMenu() {
	fmt.Println("===================================")
	fmt.Println("==        Gestión Stripe         ==")
	fmt.Println("===================================")
	fmt.Println(" 1 - Lista de productos")
	fmt.Println(" 2 - Añadir nuevo producto")
	fmt.Println(" 3 - Listar todos los precios")
	fmt.Println(" 4 - Buscar producto por ID")
	fmt.Println(" 5 - Borrar producto")
	fmt.Println("===================================")
	fmt.Print("Seleccionar opción : ")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		productsItems := products.GetAllProducts()
		views.ViewAllProducts(productsItems)
		MainMenu()
	case 2:
		forms.FormNewProduct()
	case 3:
		products.GetAllPrices()
		MainMenu()
	case 4:
		productItem := products.GetProductById(forms.FormGetProductById())
		views.ViewProduct(productItem)
		MainMenu()
	case 5:
		productItem := forms.FormDeleteProduct()
		productDeleted := products.DeleteProduct(productItem)
		views.ViewConfirmDeleted(productDeleted)
		MainMenu()
	case 6:
		products.GetProductPriseById("aa")
	}
}
