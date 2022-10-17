package routers

import(
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/4LiveStudioGit/GO_twittor/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request){
	if len(r.URL.Query().Get("pagina"))<1{
		http.Error(w,  "introduzca el parametro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil{
		http.Error(w,  "Envie el parametro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if correcto == false{
		http.Error(w,  "Error al leer los tweets", http.StatusBadRequest)
		return
	} 
		w.Header().Set("Context.Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(respuesta)
}