package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("Hello, Haasin!")
// }

// func main() {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	fmt.Println("My favorite number is", rand.Intn(10))
// }

// func main() {
// 	fmt.Printf("Now you have %v problems.\n", math.Sqrt(7))
// }

// func add(x int, y int) int {
// 	return x + y
// }

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(40, 22))
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4/9
// 	y = sum - x

// 	return
// }

// func main() {
// 	fmt.Println(split(70))
// }

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	return lim
// }

// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	fmt.Printf("x = %v\n", x)
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2*z)
// 		fmt.Printf("iteration: %v, z = %v\n", i + 1, z)
// 	}
// 	return z
// }

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	var t float64
// 	fmt.Printf("x = %v\n", x)
// 	for {
// 		z, t = z - (z*z - x) / (2*z), z
// 		fmt.Printf("z = %v, t = %v\n", z, t)
// 		if math.Abs(z-t) < 1e-6 {
// 			break
// 		}
// 	}
// 	return z
// }

// func main() {
// 	fmt.Println(Sqrt(3))
// 	fmt.Printf("Sqrt(x) = %v\n", math.Sqrt(3))
// }

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}