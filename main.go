package main

import (
	"fmt"
	"log"
	"main/controller"
	"main/template"
	"main/view"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	homeTmpl := view.NewTemplate(template.FS, "layout.html", "home.html")
	contactTmpl := view.NewTemplate(template.FS, "layout.html", "contact.html")
	faqTmpl := view.NewTemplate(template.FS, "layout.html", "faq.html")

	router.Get("/", controller.StaticHandler(homeTmpl))
	router.Get("/contact", controller.StaticHandler(contactTmpl))
	router.Get("/faq", controller.Faq(faqTmpl))

	fmt.Println("Starting the server on http://localhost:3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
