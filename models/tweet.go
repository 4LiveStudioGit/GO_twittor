package models


//Captura en el body, el mensaje que llega
type Tweet struct{

	Mensaje string `bson:"mensaje" json:"mensaje"`


}