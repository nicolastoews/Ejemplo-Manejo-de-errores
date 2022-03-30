package main

import "fmt"

func main() {
	f()
	fmt.Println("Retornando de manera normal desde f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil { //r toma el valor del recover y como es distinto de nil, continua
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
		panic(fmt.Sprintf("%v", i)) //Una vez i llegue a 4 le damos un panic y vuelve a la funcion f
	}
	defer fmt.Println("Defer en g", i)
	fmt.Println("Imprimiendo g", i)
	g(i + 1) //al llamar a la misma función dentro de si misma se genera un bucle, esta funcion se va a repetir hasta que i valga 4, por el if de mas arriba
}
