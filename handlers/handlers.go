package handlers

import (
	"github.com/4LiveStudioGit/GO_twittor/middlew"
	"github.com/4LiveStudioGit/GO_twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors" //Definir los permisos de aceso de la app
	"log"
	"net/http"
	"os"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) // EL hhtp se pone a escuchar el periekte
}
