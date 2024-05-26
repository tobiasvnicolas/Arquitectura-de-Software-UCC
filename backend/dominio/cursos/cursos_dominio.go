package dominio

type Curso struct {
	CursoID     int    `json:"id_curso"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Categoria   string `json:"categoria"`
	//FechaCreación string
	// PARA manejar fechas capaz podemos usar la librería Time (INVESTIGAR)
}
