package main

import "fmt"

func main() {
	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	a := 10
	b := 3

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Printf("a / b = %.2f\n", float64(a)/float64(b))
}
