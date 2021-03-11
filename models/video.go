package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Video es el modelo de Video de la base de MongoDB--------
type Video struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string             `bson:"nombre" json:"nombre,omitempty"`
	FechaSubida time.Time          `bson:"fechaSubida" json:"fechaSubida,omitempty"`
	Email       string             `bson:"email" json:"email"`
	Codigo      string             `bson:"codigo" json:"codigo,omitempty"`
	Link        string             `bson:"link" json:"link,omitempty"`
	Empresa     string             `bson:"empresa" json:"empresa,omitempty"`
}
