package routers

import(
	"io"
	"net/http"
	"os"
	"github.com/4LiveStudioGit/GO_twittor/bd"

)

func ObtenerAvatar(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil{
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}
	openFile, err := os.Open("uploads/avatar/"+perfil.Avatar)
	if err != nil{
		http.Error(w, "Imagen noo encontrado", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err!=nil{
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}

}