package dominio

type UsuarioData struct {
	UsuarioID    int    `json:"usuario_id"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Email        string `json:"email"`
	Passwordhash string `json:"passwordhash"`
	Tipo         string `json:"tipo"`
	//Técnicamente Nombre y Apellido no son datos del Diseño E-R
	//Sin embargo, considero apropiado agregarlos.
	//Falta trabajar con fechas. (Creación y Update)
}

type UsuariosData []UsuarioData

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	//UsuarioID int `json:"usuario_id"`
	Token string `json:"token"`
}
