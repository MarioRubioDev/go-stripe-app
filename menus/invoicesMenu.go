package menus

import "fmt"

func InvoicesMenu() {
	fmt.Println("===================================")
	fmt.Println("==          Facturas             ==")
	fmt.Println("===================================")
	fmt.Println(" 1 - Productos")
	fmt.Println(" 2 - Clientes")
	fmt.Println(" 3 - Facturas")
	fmt.Println(" -----------------------------------")
	fmt.Println(" 0 - Menú principal")
	fmt.Println("===================================")
	fmt.Print("Seleccionar opción : ")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 0:
		MainMenu()
	}

}
