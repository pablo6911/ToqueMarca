package bd

import (
	"context"
	"time"

	"github.com/pablo6911/toquemarca/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya esta en la BD-------
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("toquelike")
	col := db.Collection("usermarca")

	condicion := bson.M{"email": email}

	var resultdo models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultdo)
	ID := resultdo.ID.Hex()
	if err != nil {
		return resultdo, false, ID
	}
	return resultdo, true, ID
}

//ChequeoYaNombreVideo recibe un nombre de parametro y chequea si ya esta en la BD-------
func ChequeoYaNombreVideo(nombre string) (models.Video, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("toquelike")
	col := db.Collection("videomarca")

	condicion := bson.M{"nombre": nombre}

	var resultdo models.Video

	err := col.FindOne(ctx, condicion).Decode(&resultdo)
	ID := resultdo.ID.Hex()
	if err != nil {
		return resultdo, false, ID
	}
	return resultdo, true, ID
}
