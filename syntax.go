package main
import "fmt"
import "math"
import "time"
import "math/rand"

func main(){
	fmt.Println("Akar dari 4 adalah ", math.Sqrt(4))
	akar(16)
	jam()
}

func akar(x float64){
	fmt.Println("Akar dari ", x, "adalah ",math.Sqrt(x))
}

func jam() {
	fmt.Println("Sekarang jam:",time.Now())
	fmt.Println("Angka acak dari 1:10", rand.Intn(10))

}