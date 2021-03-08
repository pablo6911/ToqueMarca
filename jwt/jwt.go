package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pablo6911/toquemarca/models"
)

//GeneroJWT genera el encriptado con JWT-----------
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("Toque")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"fecha_nacimiento": t.FechaNacimiento,
		"ubicacion":        t.Ubicacion,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenstr, err := token.SignedString(miClave)
	if err != nil {
		return tokenstr, err
	}
	return tokenstr, nil
}
