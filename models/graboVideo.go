package models

import (
	"time"
)

//GraboVideo es el modelo de Video en la base de MongoDB--------
type GraboVideo struct {
	UserID        string    `bson:"userid" json:"userid,,omitempty"`
	NombreVideo   string    `bson:"nombreVideo" json:"nombreVideo,omitempty"`
	Fecha         time.Time `bson:"fecha" json:"fecha,omitempty"`
	Email         string    `bson:"email" json:"email"`
	Codigo        string    `bson:"codigo" json:"codigo,omitempty"`
	Link          string    `bson:"link" json:"link,omitempty"`
	NombreEmpresa string    `bson:"nombreEmpresa" json:"nombreEmpresa,omitempty"`
	Video         string    `bson:"video" json:"viedo,omitempty"`
}
