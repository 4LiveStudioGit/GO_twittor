package middlew

import (
	"net/http"
	"github.com/4LiveStudioGit/routers"
)

// Funcion para validar el JWT que nos devuelve la petición
func ValidoJWT(next htt.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		_, _, _, err := routes.ProcesarToken(r.Header.Get("Authorization"))
		if err != nil{
			http.Error(w, "Error en el token!"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHttp(w, r)
	}
}