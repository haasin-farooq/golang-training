package main

import "fmt"

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

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(40, 22))
}