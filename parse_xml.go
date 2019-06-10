package main

import "encoding/xml"
import "fmt"
import "io/ioutil"
import "net/http"

type Urlset struct {
	Locations [] Location `xml:"url"`
}

type Location struct {
	Loc string `xml:"loc"`
}

//merubah format output dari array jadi string
func (l Location) String () string {
	return fmt.Sprintf(l.Loc)
  }

func main() {
	resp, _ := http.Get("https://www.bps.go.id/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	//resp.Body.Close()
	var s Urlset
  	xml.Unmarshal(bytes, &s)
  	fmt.Println(s.Locations)
}