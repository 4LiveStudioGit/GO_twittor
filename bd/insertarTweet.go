package bd

import(
	"context"
	"time"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarTweet(tweet models.GraboTweet)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("GOTwittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userID" 	: 	tweet.UserID,
		"mensaje"	:	tweet.Mensaje,
		"fecha"		:	tweet.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil{
		return "",false,err
	}
	//Devuelve el objectID del ultimo registro insertado en la BBDD
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(),true,nil
}