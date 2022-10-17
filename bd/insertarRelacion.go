package bd

import(

	"context"
	"time"
	"github.com/4LiveStudioGit/GO_twittor/models"
)

func InsertarRelaccion(users models.Relacion)(bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Modificamos el contexto TODO y se crea un timeout a los 15 seg
	defer cancel()

	db := MongoCN.Database("GOTwittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx,users)
	if err != nil{
		return false, err
	}
	return true, nil
}