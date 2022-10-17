package bd

import(
	"context"
	"fmt"
	"time"
	"github.com/4LiveStudioGit/GO_twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	
)
/*
- ID: Usuario activo que va a realizar la busqueda
- page: numero de pagina(cada pagina mostrará 20 resultados)
- search: Palabra por la que el usuario desea filtrar la busqueda
- tipo:Tipo de busqueda
*/
func LeoUsuariosTodos(ID string, page int64, search string, tipo string)([]*models.Usuario, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("GOTwittor")
	col := db.Collection("Usuarios")

	var resultados[]*models.Usuario

	//Configuramos las opciones de la busqueda
	findOptions := options.Find()
	
	//Definimos cuantos elementos saltamos para mostrar los de la página correspondiente
	findOptions.SetSkip((page-1)*20)
	//Definimos cuantos elementos se muestran por página
	findOptions.SetLimit(20)
	
	consulta := bson.M{
		"nombre": bson.M{"$regex":`(?i)`+search},
	}

	cur, err := col.Find(ctx,consulta,findOptions)
	if err != nil{
		fmt.Println(err.Error())
		return resultados, false
	}
	var encontrado, incluir bool

	for cur.Next(ctx){
		var user models.Usuario
		err := cur.Decode(&user)
		if err !=nil{
			fmt.Println(err.Error())
			return resultados,false
		}
		var relacion models.Relacion
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID= user.ID.Hex()
		incluir = false
		encontrado, err = ConsultaRelacion(relacion)
		if tipo =="new" && encontrado ==false{
			incluir = true
		}
		if tipo =="follow" && encontrado == true{
			incluir = true
		}
		if relacion.UsuarioRelacionID == ID{
			incluir = false
		}
		if incluir == true{
			user.Password = ""
			user.Biografia = ""
			user.SitioWeb = ""
			user.Ubicacion = ""
			user.Banner = ""
			user.Email = ""

			resultados = append(resultados, &user)

		}
	}

	err = cur.Err()
		if err != nil{
			fmt.Println(err.Error())
			return resultados, false
		}
	
	cur.Close(ctx)
	return resultados,true


}