package main
import "fmt"
import "net/http"

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go berbasis web!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Belajar Go Language bersama javan")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/tentang/", about_handler)
	http.ListenAndServe(":8000", nil) 
}