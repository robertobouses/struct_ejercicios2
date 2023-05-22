package main

import "github.com/robertobouses/user"

func main() {
	var administrador1 user.Administrador

	administrador1.cargar()
	administrador1.imprimir()
}
