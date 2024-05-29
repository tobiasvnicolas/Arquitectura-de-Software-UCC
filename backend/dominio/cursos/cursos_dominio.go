package dominio

type CursoData struct {
	CursoID     int    `json:"curso_id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
