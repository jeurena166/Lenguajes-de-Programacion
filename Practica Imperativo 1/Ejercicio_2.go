package main

import (
	"fmt"
)

func imprimeFigura(tamaño int) {
	if tamaño > 0 && tamaño%2 != 0 {
		for i := 1; i <= tamaño; i += 2 {
			espacios := (tamaño - i) / 2
			for j := 0; j < espacios; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}

		for i := tamaño - 2; i >= 1; i -= 2 {
			espacios := (tamaño - i) / 2
			for j := 0; j < espacios; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	} else {
		return
	}
}

func main() {
	tamaño := 7
	imprimeFigura(tamaño)
}
