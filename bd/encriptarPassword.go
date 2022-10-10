package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {
	costo := 8 //Encriptación de 2 elevado al costo (2 elevado a 8 encriptacion de 256 veces)
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}
