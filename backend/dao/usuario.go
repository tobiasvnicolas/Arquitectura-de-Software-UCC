package dao

type Usuario struct {
	UsuarioID    int    `gorm:"primary_key;AUTO_INCREMENT"`
	Nombre       string `gorm:"type:varchar(255);not null"`
	Apellido     string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(255);not null;unique"`
	Passwordhash string `gorm:"type:varchar(255);not null"`
	Tipo         string `gorm:"type:varchar(255);not null"`
}

type Usuarios []Usuario
