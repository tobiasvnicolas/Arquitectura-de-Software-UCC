package dao

//import "time"

type Inscripcion struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	UsuarioID int `gorm:"type:int;not null"`
	CursoID   int `gorm:"type:int;not null"`
}

type Inscripciones []Inscripcion
