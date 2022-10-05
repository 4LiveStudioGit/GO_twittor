package main

import (
	"log"  //Nos permite escibir en el log
	"github.com/4livestudiogit/GO_twittor/handlers"
	"github.com/4livestudiogit/GO_twittor/bd"
)

func main(){
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}