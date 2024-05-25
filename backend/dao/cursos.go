package dao

type Curso struct {
	CursoID     int    `gorm:"primary_key;AUTO_INCREMENT"`
	Nombre      string `gorm:"type:varchar(255);not null"`
	Descripcion string `gorm:"type:varchar(255);not null"`
}

type Cursos []Curso
