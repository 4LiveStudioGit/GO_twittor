package bd

import (
	"context"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertarRegistro(user models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Modificamos el contexto TODO y se crea un timeout a los 15 seg
	defer cancel()

	db := MongoCN.Database("GOTwittor")
	col := db.Collection("Usuarios")
	pass, _ := EncriptarPassword(user.Password)
	user.Password = pass
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
