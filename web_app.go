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
	var s Urlset
	  xml.Unmarshal(bytes, &s)
	  for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}