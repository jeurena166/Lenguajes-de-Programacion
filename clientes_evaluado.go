package main

import (
	"fmt"
	"sort"
	"strings"
)

type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

func agregarCliente(listaClientes []*infoCliente, nombre, correo string, edad int32) []*infoCliente {
	cliente := &infoCliente{nombre, correo, edad}
	listaClientes = append(listaClientes, cliente)
	return listaClientes
}

func filter[T any](list []*T, f func(*T) bool) []*T {
	filtered := make([]*T, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func getApellido(nombre string) string {
	partes := strings.Split(nombre, " ")

	apellido := partes[len(partes)-1]

	apellido = strings.ToLower(apellido)

	return apellido
}

func listaClientes_ApellidoEnCorreo(listaClientes []*infoCliente) []*infoCliente {
	return filter(listaClientes, func(cliente *infoCliente) bool {
		apellido := getApellido(cliente.nombre)
		return strings.Contains(cliente.correo, apellido)
	})
}

func cantidadCorreosCostaRica(listaClientes []*infoCliente) int {
	correoCostaRica := filter(listaClientes, func(cliente *infoCliente) bool {
		return strings.HasSuffix(cliente.correo, ".cr")
	})

	return len(correoCostaRica)
}

func correosOrdenadosAlfabeticos(listaClientes []*infoCliente) []string {
	correos := make([]string, len(listaClientes))

	for i, cliente := range listaClientes {
		correos[i] = cliente.correo
	}

	sort.Strings(correos)

	return correos
}

func main() {
	var listaClientes []*infoCliente

	listaClientes = agregarCliente(listaClientes, "Oscar Viquez", "oviquez@tec.ac.cr", 44)
	listaClientes = agregarCliente(listaClientes, "Pedro Perez", "elsegundo@gmail.com", 30)
	listaClientes = agregarCliente(listaClientes, "Maria Lopez", "mlopez@hotmail.com", 18)
	listaClientes = agregarCliente(listaClientes, "Juan Rodriguez", "jrodriguez@gmail.com", 25)
	listaClientes = agregarCliente(listaClientes, "Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	listaClientes = agregarCliente(listaClientes, "Marco Rojas", "loquesea@hotmail.com", 47)
	listaClientes = agregarCliente(listaClientes, "Marta Saborio", "msaborio@gmail.com", 33)
	listaClientes = agregarCliente(listaClientes, "Camila Segura", "csegura@ice.co.cr", 19)
	listaClientes = agregarCliente(listaClientes, "Fernando Rojas", "frojas@estado.gov", 27)
	listaClientes = agregarCliente(listaClientes, "Rosa Ramirez", "risuenna@gmail.com", 50)

	for i, cliente := range listaClientes {
		fmt.Printf("Cliente %d: Name: %s, Email: %s, Age: %d\n", i+1, cliente.nombre, cliente.correo, cliente.edad)
	}

	clientesConApellidoEnCorreo := listaClientes_ApellidoEnCorreo(listaClientes)

	fmt.Println("Clientes cuyo apellido está en el correo:")
	for _, cliente := range clientesConApellidoEnCorreo {
		fmt.Printf("Nombre: %s, Correo: %s, Edad: %d\n", cliente.nombre, cliente.correo, cliente.edad)
	}

	cantidadCorreos := cantidadCorreosCostaRica(listaClientes)

	fmt.Printf("Cantidad de clientes con correo en Costa Rica: %d\n", cantidadCorreos)

	correosOrdenados := correosOrdenadosAlfabeticos(listaClientes)

	fmt.Println("Correos electrónicos ordenados alfabéticamente:")
	for _, correo := range correosOrdenados {
		fmt.Println(correo)
	}
}
