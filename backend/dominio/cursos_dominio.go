package dominio

type CursoData struct {
	CursoID    int64    `json:"curso_id"`
	Nombre       string `json:"nombre"`
	Categoria         string `json:"tipo"`
	Descripcion        string `json:"email"`
}

