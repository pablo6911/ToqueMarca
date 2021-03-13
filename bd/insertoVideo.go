package bd

import (
	"context"
	"time"

	"github.com/pablo6911/toquemarca/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoVideo graba el video como mensaje en la bd
func InsertoVideo(t models.GraboVideo) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("toquelike")
	col := db.Collection("videomarca")

	registro := bson.M{
		"userid":        t.UserID,
		"link":          t.Link,
		"codigo":        t.Codigo,
		"nombreEmpresa": t.NombreEmpresa,
		"nombreVideo":   t.NombreVideo,
		"email":         t.Email,
		"fecha":         t.Fecha,
	}
	t.Codigo, _ = EncriptarPassword(t.Codigo)

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
