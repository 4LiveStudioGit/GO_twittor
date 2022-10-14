package routers

import (
	"encoding/json"
	"net/http"
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"

)

func ModificarPerfil(w http.ResponseWriter, r *http.Request){
	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)
	if err !=nil{
		http.Error(w, "Datos Incorrectos"+ err.Error(),400)
		return
	}
	var status bool
	status, err = bd.ModificarRegistro(user, IDUsuario)
	if err !=nil{
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+ err.Error(),400)
		return
	}
	if status == false{
		http.Error(w, "No se ha modificado el registro del usuario"+ err.Error(),400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}