package jwt

import(
	"time"
	jwt "github.com/dgrijalva/jwt-go"	
	"github.com/4LiveStudioGit/GO_twittor/models"
)
//GenerarJWT genera el encriptado con JWT
func GenerarJWT(user models.Usuario)(string,error){
	miClave := []byte("minas-moria")
	payload := jwt.MapClaims{
		"email"		: 	user.Email,
		"nombre"	:	user.Nombre,
		"apellidos"	:	user.Apellidos,
		"fecha_nacimiento" :	user.FechaNacimiento,
		"_id" 				: user.ID.Hex(),
		"exp"				: time.Now().Add(time.Hour * 24).Unix()	}	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr ,err := token.SignedString(miClave)	
	if err != nil{
		return tokenStr,err
	}
	return tokenStr,nil
}