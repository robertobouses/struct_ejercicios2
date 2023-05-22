package user

type Usuario struct {
	nombre string
	clave  string
}

type Administrador struct {
	Usuario
	privilegio int
}
