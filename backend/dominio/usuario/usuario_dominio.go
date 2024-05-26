package dominio

type UsuarioData struct {
	UsuarioID    int    `json:"usuario_id"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Tipo         string `json:"tipo"`
	Email        string `json:"email"`
	Passwordhash string `json:"passwordhash"`
}

type UsuariosData []UsuarioData
