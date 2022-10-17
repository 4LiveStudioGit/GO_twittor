package bd

import (
	"context"
	"time"
	"github.com/4LiveStudioGit/GO_twittor/models"
)
func BorrarRelacion(relacion models.Relacion)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db:= MongoCN.Database("GOTwittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, relacion)
	if err != nil{
		return false, err
	}
	return true, nil
}