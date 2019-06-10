package main

import "encoding/xml"
import "fmt"
import "io/ioutil"
import "net/http"

func main() {
	// loop1()
	// loop2()
	// loop3()
	// loop4()
	// loop5()
	// loop6()
	// loop7()
	looping_xml()
}

func loop1() {
	for i:=0; i<10; i++{
		fmt.Println(i)
	}
}

func loop2() {
	i:=0
	for i<10{
		fmt.Println(i)
		i++
	}
}

func loop3(){
	//infinity loop
	for {
		fmt.Println("Do stuff.")
	}
}

func loop4(){
	x := 5
	for {
		fmt.Println("Do stuff.",x)
		x+=3
		if x > 25 {
			break
		}
	}
}

func loop5(){
	for x:=5; x<25; x+=3{
		fmt.Println("Do stuff",x)
	}
}

func loop6(){
	for x:=5; x<25; x+=3{
		fmt.Println("Do stuff",x)
	}
}

func loop7(){
	a := 3
	for x:=5; a <25; x+=3{
		fmt.Println("Do stuff",x)
		a+=4
	}
}

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

func looping_xml() {
	resp, _ := http.Get("https://www.bps.go.id/sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s Urlset
	  xml.Unmarshal(bytes, &s)
	  for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}