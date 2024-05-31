package servicios

import (
	"errors"
	"strings"

)


func Login(email string, passwordhash string) (string, error) {

	if strings.TrimSpace(email) == ""{//Si lee que no pasaste nada en email, te manda error
		return "", errors.New("email is required")
	}


}
