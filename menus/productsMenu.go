package menus

import (
	"fmt"
	"go-stripe-app/forms"
	"go-stripe-app/stripemethods/productMethods"
	"go-stripe-app/views"
)

func ProductsMenu() {
	fmt.Println("===================================")
	fmt.Println("==          Productos            ==")
	fmt.Println("===================================")
	fmt.Println(" 1 - Lista de productos")
	fmt.Println(" 2 - Añadir nuevo producto")
	fmt.Println(" 4 - Buscar producto por ID")
	fmt.Println(" ----------------------------------")
	fmt.Println(" 0 - Menú principal")
	fmt.Println("===================================")
	fmt.Print("Seleccionar opción : ")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		productsItems := productMethods.GetAllProducts()
		views.ViewAllProducts(productsItems)
		ProductsMenu()
	case 2:
		forms.FormNewProduct()
	case 3:
		productItem := productMethods.GetProductById(forms.FormGetProductById())
		views.ViewProduct(productItem)
		ProductsMenu()
	case 0:
		MainMenu()
	}
}
