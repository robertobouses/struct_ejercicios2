package user

import "fmt"

func (admin *Administrador) cargar() {
	fmt.Println("indica nombre")
	fmt.Scanln(&admin.nombre)
	fmt.Println("indica clave")
	fmt.Scanln(&admin.clave)
	fmt.Println("indica privilegio")
	fmt.Scanln(&admin.privilegio)
}

func (admin Administrador) imprimir() {

	fmt.Println("Nombre: ", admin.privilegio)
	fmt.Println("Clave: ", admin.privilegio)
	fmt.Println("Privilegio: ", admin.privilegio)

}
