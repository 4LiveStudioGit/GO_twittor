package bd

import (
	"context"
	"time"
	"fmt"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscarPerfil(ID string)(models.Usuario, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db:= MongoCN.Database("GOTwittor")
	col := db.Collection("Usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}
	

	err:= col.FindOne(ctx,condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado "+ err.Error())
		return perfil, err
	}
	return perfil, nil


}