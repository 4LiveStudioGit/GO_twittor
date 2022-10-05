package middlew

import(
	"net/http"
	"github.com/4livestudiogit/GO_twittor/bd"
)
func ChequeoBD(next http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if bd.ChequeoConnection()==0{
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}