package main
import "fmt"

func main() {
	x := 15
	a := &x         // a hanya untuk menyimpan alamat memori
	fmt.Println(a)  // cetak alamat memori tempat x disimpan
	fmt.Println(*a) // baca nilai di alamat memori a
	
	*a = 5          // memodifikasi nilai di memori a
	fmt.Println(x)  // cetak nilai baru
	*a = *a**a      // merubah nilai menjadi 5*5
	fmt.Println(x)  // cetak nilai
	fmt.Println(*a) // cetak nilai

	fmt.Println(&x) // cetak alamat memori
	fmt.Println(a)      // cetak alamat memori
}