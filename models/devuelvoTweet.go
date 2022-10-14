package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweets struct{

		ID primitive.ObjectID  `bson:"_id" json:"_id,omiempty"`
		UserID string	`bson:"userid" json:"userId,omiempty"`
		Mensaje string	`bson:"mensaje" json:"mensaje,omiempty"`
		Fecha time.Time `bson:"fecha" json:"fecha,omiempty"`
	}