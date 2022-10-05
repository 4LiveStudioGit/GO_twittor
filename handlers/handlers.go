package handlers

import(
	"log"
	"net/http"
	"os"
	"github.com/4livestudiogit/GO_twittor/middlew"
	"github.com/4livestudiogit/GO_twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"  //Definir los permisos de aceso de la app
)

func Manejadores(){
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT =="" {
		PORT="8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT,handler)) // EL hhtp se pone a escuchar el periekte
}