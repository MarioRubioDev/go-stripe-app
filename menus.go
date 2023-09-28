package main

import "fmt"

func mainMenu() {
	fmt.Println("===================================")
	fmt.Println("==        Gestión Stripe         ==")
	fmt.Println("===================================")
	fmt.Println("1 - Lista de productos")
	fmt.Println("2 - Añadir nuevo producto")
	fmt.Println("3 - Listar todos los precios")
	fmt.Println("4 - Buscar producto por ID")
	fmt.Println("5 - Borrar producto")
	fmt.Println("===================================")
	fmt.Print("Seleccionar opción : ")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		products := getAllProducts()
		viewAllProducts(products)
		mainMenu()
	case 2:
		formNewProduct()
	case 3:
		getAllPrices()
	case 4:
		product := getProductById(formGetProductById())
		viewProduct(product)
		mainMenu()
	case 5:
		product := formDeleteProduct()
		productDeleted := deleteProduct(product)
		viewConfirmDeleted(productDeleted)
		mainMenu()
	}
}
