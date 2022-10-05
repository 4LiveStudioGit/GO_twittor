package routers

import(
	"encoding/json"
	"net/http"
	"github.com/4livestudiogit/GO_twittor/bd"
	"github.com/4livestudiogit/GO_twittor/models"

)

func Registro(w http.ResponseWriter, r *http.Request){

		var user models.Usuario
		err := json.NewDecoder(r.Body).Decode(&user) // Una vez se lee el Body se destruye, por eso tenemos que volcarlo en el puntero user
		if err !=nil{
			http.Error(w, "Error en los datos recibidos" + err.Error(), 400 )
			return
		}
		if len(user.Email) == 0{
			http.Error(w, "El email de usuario es necesario", 400)
			return
		}
		if len(user.Password) < 6{
			http.Error(w, "La password al menos necesita 6 caracteres", 400)
			return
		}

		_, encontrado, _ := bd.ChequeoYaExisteUsuario(user.Email)
		if encontrado == true{
			http.Error(w, "Ya existe un usuario asociado a este E-Mail", 400)
			return
		}
		_, status, err:= bd.InsertarRegistro(user)
		if err !=nil{
			http.Error(w, "Ocurrio un error y no se puedo almacenar el registro"+err.Error(), 400)
			return
		}

		if status == false{
			http.Error(w, "No se puedo almacenar el registro", 400)
			return
		}
		w.WriteHeader(http.StatusCreated)
}