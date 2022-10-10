package routers

import (
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"
)

//Variables devueltas por el modelo visibles por todos los ENDPOINTS
var Email string
var IDUsuario string
//Procesa el token para extraer sus valores
func ProcesarToken(tk string)(*models.Claim, bool, string, error){
miClave :=[]byte("Las minas de moria estan en el señor de los anillos")
claims := &models.Claim{}

splitToken := strings.Split(tk,"Bearer")
if len(splitToken) != 2{
	return claims,false,string(""), errors.New("formato de token invalido")
}

//Eliminamos espacion en blanco
tk = strings.TrimSpace(splitToken[1])

tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
	return miClave, nil
})
if err == nil{
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
	if encontrado == true{
			Email  = claims.Email
			IDUsuario = claims.ID.Hex()
	}
	return claims,encontrado, IDUsuario,nil
}
if !tkn.Valid{
	return claims, false, string(""), errors.New("token invaldo")
}
return claims, false, string(""), err

}