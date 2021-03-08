package bd

import (
	"context"
	"time"

	"github.com/pablo6911/toquemarca/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertarRegistroVideo es la parada final con la BD para insertar los datos del user-----
func InsertarRegistroVideo(u models.Video) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("toquelike")
	col := db.Collection("videomarca")

	u.Codigo, _ = EncriptarPassword(u.Codigo)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	//pasando objectId a string
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
