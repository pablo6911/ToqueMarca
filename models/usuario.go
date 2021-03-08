package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario es el modelo de usuario de la base de MongoDB--------
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
}

//Video es el modelo de Video de la base de MongoDB--------
type Video struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string             `bson:"nombre" json:"nombre,omitempty"`
	FechaSubida time.Time          `bson:"fechaSubida" json:"fechaSubida,omitempty"`
	Email       string             `bson:"email" json:"email"`
	Codigo      string             `bson:"password" json:"password,omitempty"`
	Link        string             `bson:"link" json:"link,omitempty"`
	Empresa     string             `bson:"empresa" json:"empresa,omitempty"`
}
