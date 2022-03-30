package main

import "fmt"

func main() {
	f()
	fmt.Println("Retornando de manera normal desde f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperando en f", r)
		}
	}()
	fmt.Println("Llamando g.")
	g(0)
	fmt.Println("Retornando normalmente desde g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer en g", i)
	fmt.Println("Imprimiendo g", i)
	g(i + 1)
}
