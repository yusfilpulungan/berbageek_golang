package main
import "fmt"

type car struct {
	gas_pedal uint16     //nilai minimal: 0,      maksimal: 65535    16bit
	brake_pedal uint16   //nilai minimal: 0,      maksimal: 65535
	steering_wheel int16 //nilai minimal: -32768  maksimal: 32768
	top_speed_kmh float64 
}

func main() {
	//a_car := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
	a_car := car{22314,0,12562,225.0}
	fmt.Println("gas_pedal:",a_car.gas_pedal, "brake_pedal:",a_car.brake_pedal,"steering_wheel:",a_car.steering_wheel)
}