package main

import "fmt"
import "io/ioutil"
import "net/http"

func main() {
	resp, _ := http.Get("https://www.bps.go.id/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}