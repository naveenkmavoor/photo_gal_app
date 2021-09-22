package main

import ( 
	"net/http"

	"github.com/gorilla/mux"
)
 
 

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
 
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
 
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	 
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	 
}
 
func main() { 
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)
	router.NotFoundHandler = http.HandlerFunc(notFound) 
	http.ListenAndServe(":8082", router)
}
