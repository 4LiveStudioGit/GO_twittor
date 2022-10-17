package routers

import(
	"net/http"
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"

)

func AltaRelacion(w http.ResponseWriter, r *http.Request){

	
	//Por url se envia el id de la persona a la que quiere seguir, la id del usuario la tenemos en una variable global IDUsuario
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w,"Es necesario el ID del usuario", http.StatusBadRequest )
		return
	}

	var users models.Relacion
	users.UsuarioID = IDUsuario
	users.UsuarioRelacionID = ID

	status, err := bd.InsertarRelaccion(users)
	if err != nil{
		http.Error(w,"Ocurrio un error al insertar la relacion "+err.Error(), http.StatusBadRequest )
		return
	}
	if status == false{
		http.Error(w,"NO se ha logrado insertar la realaccion "+err.Error(), http.StatusBadRequest )
		return
	}
	w.WriteHeader(http.StatusCreated)
}