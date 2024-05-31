package modelo

type Curso struct {
	CursoID     int64  `gorm:"primary_key;AUTO_INCREMENT"`
	Nombre      string `gorm:"type:varchar(255);not null"`
	Descripcion string `gorm:"type:varchar(255);not null"`
	Categoria   string `gorm:"type:varchar(255);not null"`
}

type Cursos []Curso

/*
import "time"

type Course struct {
	ID           int64     // Course ID
	Title        string    // Course title
	Description  string    // Course description
	Category     string    // Course Category. Allowed values: to be defined
	ImageURL     string    // Course image URL
	CreationDate time.Time // Course creation date
	LastUpdated  time.Time // Course last updated date
}
*/
