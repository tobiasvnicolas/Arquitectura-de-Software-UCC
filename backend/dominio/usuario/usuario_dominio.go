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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Resultado struct {
	Mensaje string `json:"mensaje"`
}
