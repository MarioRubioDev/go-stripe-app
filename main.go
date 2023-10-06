package main

import (
	"fmt"
	"go-stripe-app/menus"
	"html/template"
	"log"
	"net/http"
	"os"
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
		http.HandleFunc("/customers", customers)
		http.HandleFunc("/invoices", invoices)
		fmt.Println("Listening al localhost:8080 ...")
		http.ListenAndServe(":8080", nil)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func products(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "products.html", nil)
}
func customers(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "customers.html", nil)
}
func invoices(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "invoices.html", nil)
}
