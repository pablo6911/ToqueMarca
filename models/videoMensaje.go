package models

//VideoMensaje captura del body el mensaje que nos llega
type VideoMensaje struct {
	NombreVideo   string `bson:"nombreVideo" json:"nombreVideo"`
	Email         string `bson:"email" json:"email"`
	Codigo        string `bson:"codigo" json:"codigo"`
	Link          string `bson:"link" json:"link"`
	NombreEmpresa string `bson:"nombreEmpresa" json:"nombreEmpresa"`
}
