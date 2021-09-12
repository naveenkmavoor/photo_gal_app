package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch please sent an email to <a href=\"shootmequest@gmail.com\">shootmequest.com</a>.")

}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Questions asked before</h1><br><h3>What does it mean to be an software engineer by passion</h3>")

}
 
func main() {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", home)
	mux.HandleFunc("/faq", faq)
	mux.HandleFunc("/contact", contact) 
	http.NotFoundHandler()
	http.ListenAndServe(":8082", mux)
}
