package routers

import(
	"net/http"
	"github.com/4LiveStudioGit/GO_twittor/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err:= bd.BorrarTweet(ID, IDUsuario)
	if err !=nil{
		http.Error(w, "Ocurrio un error al eliminar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}