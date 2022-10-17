package routers

import(
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"
	

)

func SubirBanner(w http.ResponseWriter, r *http.Request){
	file, handler,err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename,".")[1]
	var archivo string = "uploads/Banners/" + IDUsuario + "." + extension
	f, err:= os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		http.Error(w, "Error al subir la imagem!!!!"+err.Error(),http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil{
		http.Error(w, "Error al copiar la imagem!!!!"+err.Error(),http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificarRegistro(usuario, IDUsuario)
	if err != nil || status == false{
		http.Error(w, "Error al copiar el banner a la BD!!!!"+err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader (http.StatusCreated)
}