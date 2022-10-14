package bd

import(
	"context"
	"time"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func ModificarRegistro(user models.Usuario, ID string)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("GOTwittor")
	col := db.Collection("Usuarios")

	registro := make(map[string]interface{})
	/*
	Chequeamos si los valores que vienen tiene valores (quiere decir modificaciones)
	o si vienen vacios quiere decir que no se modifican
	*/
	if len(user.Nombre) > 0 {
		registro ["nombre"] = user.Nombre
	}

	if len(user.Apellidos) > 0 {
		registro ["apellidos"] = user.Apellidos
	}
	
		registro ["fechaNacimiento"] = user.FechaNacimiento

	if len(user.Email) > 0 {
		registro ["email"] = user.Email
	}
	if len(user.Password) > 0 {
		registro ["passwords"] = user.Password
	}
	if len(user.Avatar) > 0 {
		registro ["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registro ["banner"] = user.Banner	
	}
	if len(user.Biografia) > 0 {
		registro ["biografia"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		registro ["ubicacion"] = user.Ubicacion
	}
	if len(user.SitioWeb) > 0 {
		registro ["sitioWeb"] = user.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_,err := col.UpdateOne(ctx,filtro,updtString)
	if err != nil{
		return false, err
	}
	return true, nil
}