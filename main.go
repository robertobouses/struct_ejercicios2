package main

import "github.com/robertobouses/struct_ejercicios2/user"

func main() {
	var administrador1 user.Administrador

	administrador1.Cargar()
	administrador1.Imprimir()

	var usuario1 user.Usuario
	usuario1.CargarUsuario()
	usuario1.ImprimirUsuario()

}
