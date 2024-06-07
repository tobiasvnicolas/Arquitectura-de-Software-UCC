package dominio

type InscripcionData struct {
	ID        int64 `json:"inscripcion_id"`
	UsuarioID int64 `json:"usuario_id"`
	CursoID   int64 `json:"curso_id"`
}

type InscripcionesData []InscripcionData

type InscripcionRequest struct {
	UserID   int64 `json:"usuario_id"`
	CourseID int64 `json:"curso_id"`
}