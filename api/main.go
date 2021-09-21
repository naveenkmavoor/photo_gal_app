package main

import (
	"html/template"
	viewtemp "main/views"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template
var (
	homeView    *viewtemp.View
	contactView *viewtemp.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
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
	homeView = viewtemp.NewView("views/home.gohtml")
	contactView = viewtemp.NewView("views/contact.gohtml")
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)
	router.NotFoundHandler = http.HandlerFunc(notFound) 
	http.ListenAndServe(":8082", router)
}
