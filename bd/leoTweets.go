package bd

import(
	"context"
	"time"
	"log"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64)([]*models.DevuelvoTweets,bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("GOTwittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userID":ID,
	}

	
	opciones := options.Find()
	//Cuantos tweets trae cada vez
	opciones.SetLimit(20)
	//Ordenamos por el campo fecha en formato descendente
	opciones.SetSort(bson.D{{Key:"fecha", Value: -1}})
	//Indicamos la pagina en la que esta y salta los tweets anteriores a la página
	opciones.SetSkip((pagina-1)*20)
	
	//cursor es un puntero donde se almacenan los resultados
	cursor, err := col.Find(ctx, condicion, opciones)
	if err!= nil{
		log.Fatal(err.Error())
		return resultados, false
	}
	
	for cursor.Next(context.TODO()){
		var registro models.DevuelvoTweets
		err = cursor.Decode(&registro)
		if err!= nil{
			return resultados,false
		}
		
		resultados = append(resultados, &registro)

	}
	return resultados,true
}