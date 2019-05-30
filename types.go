package main
import "fmt"

//bool
//string
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte ini alias untuk uint8
//rune ini alias untuk int32 yang merepresentasikan sebuah Unicode code point
//float32 float64
//complex64 complex128

func main() {
	var angka1 float64 = 5.6
	var angka2 float64 = 9.5
	
	angka3, angka4 := 5.0, 10.5

    fmt.Println("Penjumlahan", angka1, "+", angka2, "=", tambah(angka1,angka2))
	fmt.Println("Penjumlahan", angka3, "+", angka4, "=", tambah(angka3, angka4))
}

func tambah(x float64, y float64) float64 {
	return x+y
}

