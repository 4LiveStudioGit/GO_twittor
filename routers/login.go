package routers

import (

	"encoding/json"
	"net/http"
	"time"
	

	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"github.com/4LiveStudioGit/GO_twittor/jwt"
)

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")

	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)
	if err !=nil{
		http.Error(w, "Usuario o Contraseña erroneos"+err.Error(),400)
		return
	}

	if len(user.Email) == 0{
		http.Error(w, "ES necesario introducir un Email",400)
		return
	}

	documento, existe := bd.IntentoLogin(user.Email, user.Password)
	if existe == false{
		http.Error(w, "Usuario o Contraseña erroneos",400)
		return
	}

	jwtKey , err := jwt.GenerarJWT(documento) //jwtKey variable de un jason web token
	if err != nil{
		http.Error(w, "Ocurrio un error al intentar generar el token"+err.Error(),400)
		return
	}
	resp:= models.RespuestaLogin{
		Token:jwtKey,
	}

	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

expirationTime := time.Now().Add(24*time.Hour)
http.SetCookie(w, &http.Cookie{
	Name: "token",
	Value: jwtKey,
	Expires: expirationTime,
})

}