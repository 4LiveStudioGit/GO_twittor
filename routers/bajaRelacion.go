package routers

import(
	"net/http"
	"fmt"
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID
	fmt.Println(relacion.UsuarioID)
	fmt.Println(relacion.UsuarioRelacionID)


	status, err := bd.BorrarRelacion(relacion)
	if err != nil{
		http.Error(w,"Ocurrio un error al borrar la relacion "+err.Error(), http.StatusBadRequest )
		return
	}
	if status == false{
		http.Error(w,"NO se ha logrado borrar la realaccion "+err.Error(), http.StatusBadRequest )
		return
	}
	w.WriteHeader(http.StatusCreated)
}