package dominio

type UsuarioData struct {
	UsuarioID    int    `json:"usuario_id"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Email        string `json:"email"`
	Passwordhash string `json:"passwordhash"`
	Tipo         string `json:"tipo"`
}

type UsuariosData []UsuarioData
