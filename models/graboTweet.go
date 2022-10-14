package models

import "time"

type GraboTweet struct{
	UserID string `bson:"userid" json:"userid,omiempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omiempty"`
	Fecha time.Time `bson:"fecha" json:"fecha,omiempty"`

}