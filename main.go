package main

import (
	"fmt"
	"go-stripe-app/menus"
	"go-stripe-app/stripemethods/productMethods"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var tpl *template.Template

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "t" {
		_, err := template.ParseGlob("txtTemplates/*")
		if err != nil {
			log.Fatalln(err)
		}
		menus.MainMenu()
		os.Exit(0)
	}

	if args[0] == "w" {

		tpl = template.Must(template.ParseGlob("templates/*.html"))
		http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets"))))

		http.HandleFunc("/", index)
		http.HandleFunc("/products", products)
		http.HandleFunc("/addproduct", addProduct)
		http.HandleFunc("/allproducts", allProducts)
		http.HandleFunc("/customers", customers)
		http.HandleFunc("/invoices", invoices)
		fmt.Println("Listening al localhost:8080 ...")
		fmt.Println("Run server: http://localhost:8080")
		http.ListenAndServe(":8080", nil)
	}

}

func addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		priceConvert, err := strconv.ParseFloat(r.FormValue("price"), 10)
		if err != nil {
			log.Println(err)
		}

		newProduct := productMethods.StripeProduct{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			PriceValue:  int64(priceConvert * 100),
		}

		newProductID, newPriceID := productMethods.AddNewStripeProduct(newProduct)

		fmt.Println("Código producto : ", newProductID)
		fmt.Println("Código precio : ", newPriceID)

		tpl.ExecuteTemplate(w, "products.html", nil)

	} else {
		tpl.ExecuteTemplate(w, "addproduct.html", nil)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
}

func products(w http.ResponseWriter, r *http.Request) {
	productList := productMethods.GetAllProducts()
	tpl.ExecuteTemplate(w, "products.html", productList)
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	productList := productMethods.GetAllProducts()

	productListItems := []productMethods.StripeProduct{}

	for productList.Next() {
		productListItems = append(
			productListItems,
			productMethods.StripeProduct{
				Name:        productList.Product().Name,
				Description: productList.Product().Description,
			})

		fmt.Println(productList.Product().Name)
	}

	tpl.ExecuteTemplate(w, "allproducts.html", productListItems)
}

func customers(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "customers.html", nil)
}

func invoices(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "invoices.html", nil)
}
