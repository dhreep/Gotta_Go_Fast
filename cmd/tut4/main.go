package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (0.5*a*t*t + v0*t + s0)
	}
}

func main() {
	var a, v0, s0, t float64
	fmt.Println("Enter the values for:")
	fmt.Print("Acceleration: ")
	fmt.Scanf("%f\n", &a)
	fmt.Print("Initial Velocity: ")
	fmt.Scanf("%f\n", &v0)
	fmt.Print("Intial Displacement: ")
	fmt.Scanf("%f\n", &s0)
	fm := GenDisplaceFn(a, v0, s0)
	fmt.Print("Enter the time: ")
	fmt.Scanf("%f\n", &t)
	fmt.Println("The total distance is:", fm(t))

}
