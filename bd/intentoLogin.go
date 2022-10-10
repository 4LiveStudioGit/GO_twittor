package bd

import (
	"github.com/4LiveStudioGit/GO_twittor/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	user, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
