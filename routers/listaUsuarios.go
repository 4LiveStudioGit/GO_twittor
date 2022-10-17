package routers

import(
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/4LiveStudioGit/GO_twittor/bd"

)

func ListaUsuarios( w http.ResponseWriter, r *http.Request){
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page) //Convierte un alfabetico a integer
	if err != nil{
		http.Error(w, "Debe enviar el parametro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag :=  int64(pagTemp)

	resultado, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser )
	if status == false{
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Context.Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}