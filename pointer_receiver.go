package main
import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 // min 0 to max 65535
	brake_pedal uint16
	steering_wheel uint16 // -32k to +32k
	top_speed_kmh float64
}

// methods are not same as function since they are associated with structures
// return kmh() in float
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

// Should I define methods on values or pointers?
// https://golang.org/doc/faq#methods_on_values_or_pointers

func (c *car) mph() float64 {
	c.top_speed_kmh = 500
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{gas_pedal: 65000, 
				brake_pedal: 0, 
				steering_wheel: 12361,
				top_speed_kmh: 225.0}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}