package dao

type Usuario struct {
	UsuarioID    int    `gorm:"primary_key;AUTO_INCREMENT; not null"`
	Nombre       string `gorm:"type:varchar(255)"`
	Apellido     string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(255)"`
	Passwordhash string `gorm:"type:varchar(255)"`
	Tipo         string `gorm:"type:varchar(255)"`
}

type Usuarios []Usuario

/*
type User struct {
	ID           int64     // User ID
	Email        string    // User Email
	PasswordHash string    // User Password Hash
	Type         string    // User Type. Allowed values: admin, normal
	CreationDate time.Time // User creation date
	LastUpdated  time.Time // User last updated date
}


package dao

type Usuario struct {
	UsuarioID    int    `gorm:"primary_key;AUTO_INCREMENT; not null"`
	Nombre       string `gorm:"type:varchar(255)"`
	Apellido     string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(255)"`
	Passwordhash string `gorm:"type:varchar(255)"`
	Tipo         string `gorm:"type:varchar(255)"`
}

type Usuarios []Usuario
*/
