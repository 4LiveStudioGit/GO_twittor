package routers

import (
	
	"encoding/json"
	"net/http"
	"github.com/4LiveStudioGit/GO_twittor/bd"
)


func VerPerfil(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscarPerfil(ID)
	if err != nil{
		http.Error(w, "Ocurrió un error al buscar el registro " + err.Error() , 400)
	}

	w.Header().Set("context-type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
	
}