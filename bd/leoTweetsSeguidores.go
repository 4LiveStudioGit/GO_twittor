package bd

import(
	"context"
	"time"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores( ID string, pagina int)([]models.DevuelvoTweetsSeguidores, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("GOTwittor")
	col := db.Collection("relacion")

	skip := (pagina-1)*20

	condiciones := make([]bson.M,0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid":ID}})
	condiciones = append(condiciones, bson.M{
		//Se utiliza para enlazar dos tablas
		"$lookup": bson.M{
			"from": "tweet",
			"localField": "usuariorelacionid",
			"foreignField": "userid",
			"as": "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind":"$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha":-1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var resultado []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &resultado)
	if err != nil{
		return resultado, false
	}
	return resultado, true

}
