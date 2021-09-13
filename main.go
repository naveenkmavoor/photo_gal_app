package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tpl.ExecuteTemplate(w, "home.gohtml", nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tpl.ExecuteTemplate(w, "contact.gohtml", nil); err != nil {
		panic(err)
	}
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tpl.ExecuteTemplate(w, "faq.gohtml", nil); err != nil {
		panic(err)
	}
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	if err := tpl.ExecuteTemplate(w, "pagenotfound.gohtml", nil); err != nil {
		panic(err)
	}
}
func init() {
	tpl = template.Must(template.ParseGlob("views/*.gohtml"))

}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8082", router)
}
