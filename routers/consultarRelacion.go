package routers

import(
	"net/http"
	"encoding/json"
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"
)

func ConsultarRelacion(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultaRelacion(relacion)
	if err != nil || status == false{
		resp.Status = false
	}else{
		resp.Status = true
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)


}