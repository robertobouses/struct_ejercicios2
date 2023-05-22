package user

import "fmt"

func (admin *Administrador) Cargar() {
	fmt.Println("indica nombre del admin")
	fmt.Scanln(&admin.nombre)
	fmt.Println("indica clave del admin")
	fmt.Scanln(&admin.clave)
	fmt.Println("indica privilegio del admin")
	fmt.Scanln(&admin.privilegio)
}

func (admin Administrador) Imprimir() {

	fmt.Println("Nombre del admin: ", admin.nombre)
	fmt.Println("Clave del admin: ", admin.clave)
	fmt.Println("Privilegio del admin: ", admin.privilegio)
	fmt.Println("---------------------------------------")
}
