package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("log.txt") //creamos el archivo log
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()  //siempre cerrar el archivo abierto, se usa defer para que lo cierre al final del programa
	log.SetOutput(f) //Indicamos que el log se guarda en el archivo que indica la variable f, que es log.txt

	f2, err := os.Open("noexiste.txt")
	if err != nil {
		log.Println("Ocurrió un error:", err)
	}
	defer f2.Close()
	fmt.Println("Revisa el archivo log.txt en el directorio de este programa.")
}
