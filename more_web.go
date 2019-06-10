package main

import "fmt"
import "net/http"

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go itu Rapi!")
	fmt.Fprintln(w, "Go itu Cepat!")
	fmt.Fprintln(w, "dan Sederhana!")
	fmt.Fprintf(w, "<p>Kamu %s menambahkan %s</p>", "bisa", "<strong>variabel</strong>")
	fmt.Fprintf(w, `
		<h6>Kamu juga dapat melakukan</h6>
		<h5>banyak baris</h5>
		<h4>dalam satu %s</h4>`, "formatted print")
}

func tentang_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ahli Desain Web oleh Harrison Kinsley")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/tentang/", tentang_handler)
	http.ListenAndServe(":8000", nil) 
}