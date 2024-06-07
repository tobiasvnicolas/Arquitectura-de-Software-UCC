package dao

type Inscripcion struct {
	ID        int64 `gorm:"primary_key;AUTO_INCREMENT"`
	UsuarioID int64 `gorm:"type:int;not null"`
	CursoID   int64 `gorm:"type:int;not null"`
}

type Inscripciones []Inscripcion
