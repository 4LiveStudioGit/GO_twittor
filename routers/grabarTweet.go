package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/models"
)

func GrabarTweet(w http.ResponseWriter, r *http.Request){
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID :IDUsuario,
		Mensaje : mensaje.Mensaje,
		Fecha : time.Now(),
	}
	_, status, err := bd.InsertarTweet(registro)
	if err != nil{
		http.Error(w, "Ocurrio un error al intentar insertar el tweet"+err.Error(),400)
		return
	}
	if status == false{
		http.Error(w, "El tweet no se ha podido almacenar",400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}