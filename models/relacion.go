package models

type Relacion struct{
	//ID del Usuario Actual
	UsuarioID 			string 		`bson:"usuarioid" json:"usuarioId"`
	//ID del Usuario al que desea seguir
	UsuarioRelacionID 	string		`bson:"usuariorelacionid" json:"usuarioRelaccionId"`

}