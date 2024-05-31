package dominio

type CursoData struct {
	CursoID     int64    `json:"curso_id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Categoria 	string `json:"categoria"`
}
