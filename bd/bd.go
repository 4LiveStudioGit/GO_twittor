package bd

import(
	"context"
	"log" //Nos permite manejar el log del programa
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)
var MongoCN=conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://belerian:Ironhammer1@cluster0.r6ykvuv.mongodb.net/?retryWrites=true&w=majority")


/*Función que permite conectar a la BBDD*/
func conectarBD() *mongo.Client{
	client,err := mongo.Connect(context.TODO(),clientOptions) //Realizamos la conexion con mongo. "TODO" Sin restricciones de contexto
	if err != nil{
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(),nil)
	if err != nil{
		log.Fatal(err)
		return client
	}

	log.Println("Conexión exitosa con la BD")
	return client

}

func ChequeoConnection() int{
	err := MongoCN.Ping(context.TODO(),nil)
	if err != nil{
		return  0
	}
	return 1

}