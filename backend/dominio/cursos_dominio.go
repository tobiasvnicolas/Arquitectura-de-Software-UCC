package dominio

type CursoData struct {
	CursoID    int64    `json:"curso_id"`
	Nombre       string `json:"nombre"`
	Categoria         string `json:"categoria"`
	Descripcion        string `json:"descripcion"`
}

type CursoBuscado  struct{
	Results []CursoData `json:"results"`
}