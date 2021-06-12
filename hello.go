package main

import (
	"fmt"
	"math"
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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}