package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16     //nilai minimal: 0,      maksimal: 65535    16bit
	brake_pedal uint16   //nilai minimal: 0,      maksimal: 65535
	steering_wheel int16 //nilai minimal: -32768  maksimal: 32768
	top_speed_kmh float64 
}

//kilometer per jam
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax) //top speed: 225. So do pedal * (225/65535)
}

//mil per jam
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple/usixteenbitmax)  //top speed: 140 mph.So do pedal * (225/1.60934/65535)
}

//Memodifikasi struct dari pointer type
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	//a_car := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562}
	a_car := car{22314,0,12562,225.0}
	fmt.Println("gas_pedal:",a_car.gas_pedal, "brake_pedal:",a_car.brake_pedal,"steering_wheel:",a_car.steering_wheel)
	fmt.Println("Car is going:",a_car.mph(),"MPH, ",a_car.kmh(),"KMH, Top Speed:",a_car.top_speed_kmh)
	a_car.new_top_speed(500)
	fmt.Println("Car is going:",a_car.mph(),"MPH, ",a_car.kmh(),"KMH, Top Speed:",a_car.top_speed_kmh)
}