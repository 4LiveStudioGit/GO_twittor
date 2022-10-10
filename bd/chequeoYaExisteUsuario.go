package bd

import (
	"context"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("GOTwittor")
	col := db.Collection("Usuarios")

	condicion := bson.M{"email": email} //creamos la condicion
	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado) //Si no encuentra nada devuelve error sino lo almacena en resultado
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
