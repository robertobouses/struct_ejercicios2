package user

import "fmt"

func (usu *Usuario) CargarUsuario() {
	fmt.Println("indica nombre del usuario")
	fmt.Scanln(&usu.nombre)
	fmt.Println("indica clave del usuario")
	fmt.Scanln(&usu.clave)
}

func (usu Usuario) ImprimirUsuario() {
	fmt.Println("el nombre del usuario es", usu.nombre)
	fmt.Println("la clave del usuario es", usu.clave)
	fmt.Println("datos completos usuario", usu)
	fmt.Println("---------------------------------------")
}
