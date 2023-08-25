package main

import (
	"fmt"
)

func rotacion(arr []string, pos int, direccion int) []string {
	tamaño := len(arr)
	pos = pos % tamaño

	if pos < 0 {
		pos = tamaño + pos
	}

	if direccion == 0 { // Rotación a la izquierda
		return append(arr[pos:], arr[:pos]...)
	} else if direccion == 1 { // Rotación a la derecha
		return append(arr[tamaño-pos:], arr[:tamaño-pos]...)
	}

	return arr
}

func main() {
	secuencia := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	posiciones := 4
	direccion := 0 // 0 para izquierda, 1 para derecha

	rot := rotacion(secuencia, posiciones, direccion)
	fmt.Println("Secuencia rotada:", rot)
}
