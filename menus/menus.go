package menus

import (
	"fmt"
)

func MainMenu() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")
	fmt.Println("===================================")
	fmt.Println("==        Gestión Stripe         ==")
	fmt.Println("===================================")
	fmt.Println(" 1 - Productos")
	fmt.Println(" 2 - Clientes")
	fmt.Println(" 3 - Facturas")
	fmt.Println("===================================")
	fmt.Print("Seleccionar opción : ")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		fmt.Printf("\x1bc")
		ProductsMenu()
	case 2:
		fmt.Printf("\x1bc")
		CustomersMenu()
	case 3:
		fmt.Printf("\x1bc")
		InvoicesMenu()
	}
}
