package main

import (
	"fmt"
	"strings"
)

func calculoVariables(texto string) [3]int {
	var numCaracteres, numPalabras, numLineas int = 0, 0, 0

	lineas := strings.Split(texto, "\n")
	numLineas = len(lineas)

	palabras := strings.Fields(texto)
	numPalabras = len(palabras)

	numCaracteres = len(texto)

	resultados := [3]int{numCaracteres, numPalabras, numLineas}
	return resultados
}

func main() {

	texto1 := "Hola Mundo!"

	res1 := calculoVariables(texto1)

	fmt.Println(texto1)
	fmt.Printf("Numero de Caracteres: %d\n", res1[0])
	fmt.Printf("Numero de Palabras: %d\n", res1[1])
	fmt.Printf("Numero de Lineas: %d\n", res1[2])
	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

	texto2 := "Hola Mundo!\nSoy tu computadora\nTe estoy hablando!"

	res2 := calculoVariables(texto2)

	fmt.Println(texto2)
	fmt.Printf("Numero de Caracteres: %d\n", res2[0])
	fmt.Printf("Numero de Palabras: %d\n", res2[1])
	fmt.Printf("Numero de Lineas: %d\n", res2[2])

}
