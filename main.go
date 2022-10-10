package main

import (
	
	"github.com/4LiveStudioGit/GO_twittor/bd"
	"github.com/4LiveStudioGit/GO_twittor/handlers"
	"log" //Nos permite escibir en el log
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
