package bd

import(
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/4LiveStudioGit/GO_twittor/models"
)

func ConsultaRelacion(relacion models.Relacion)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db:= MongoCN.Database("GOTwittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid": relacion.UsuarioID,
		"usuariorelacionid": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	fmt.Println(resultado)
	if err != nil{
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil


}