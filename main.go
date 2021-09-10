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

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8082", nil)
}
