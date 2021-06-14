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

// func main() {
// 	i, j := 42, 2701

// 	p := &i         // point to i
// 	fmt.Println(*p) // read i through the pointer
// 	*p = 21         // set i through the pointer
// 	fmt.Println(i)  // see the new value of i

// 	p = &j         // point to j
// 	*p = *p / 37   // divide j through the pointer
// 	fmt.Println(j) // see the new value of j
// }

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.X = 1e9
// 	fmt.Println(v)
// }

// func main() {
// 	q := []int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(q)

// 	r := []bool{true, true, false}
// 	fmt.Println(r)

// 	s := []struct{
// 		i int
// 		b bool
// 	}{
// 		{2, true},
// 		{3, false},
// 		{5, true},
// 		{7, true},
// 		{11, false},
// 		{13, true},
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	pow := make([]int, 10)

// 	for i := range pow {
// 		pow[i] = 1 << uint(i)
// 	}
// 	for _, v := range pow {
// 		fmt.Println(v)
// 	}
// }

// func Pic(dx, dy int) [][]uint8 {
// 	s := make([][]uint8, dy)
// 	for i := range s {
// 		s[i] = make([]uint8, dx)
// 	}
// 	return s
// }

// func main() {
// 	pic.Show(Pic)
// }

// func WordCount(s string) map[string]int {
// 	m := make(map[string]int)
// 	f := strings.Fields(s)
// 	for _, v := range f {
// 		m[v] = 0
// 	}
// 	for _, v := range f {
// 		m[v] = m[v] + 1
// 	}
// 	return m
// }

// func main() {
// 	wc.Test(WordCount)
// }

// // fibonacci is a function that returns
// // a function that returns an int.
// func fibonacci() func() int {
// 	prev := 0
// 	num := 1
// 	return func() int {
// 		temp := num
// 		num += prev
// 		prev = temp
// 		return num
// 	}
// }

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Printf("v scaled by 10 = %v\n", v);
	fmt.Printf("Absolute value of v = %v\n", v.Abs());
}