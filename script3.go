package main

import (
	"fmt"
)

func main() {
	go maFonction()
	fmt.Println("Fin du programme")
}

func maFonction() {
	fmt.Println("j'ai fini !")
}
