package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pablo6911/toquemarca/bd"
	"github.com/pablo6911/toquemarca/jwt"
	"github.com/pablo6911/toquemarca/models"
)

//Registro es la func para crear en la BD el registro de user-------
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requqerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Desbes especificar una contraseña de almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un Usuario con ese Email", 400)
		return
	}
	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un Error al Resgistrar el Usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//Video es la func para crear en la BD el registro de user-------
func Video(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Video

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de de la empresa es requqerido", 400)
		return
	}
	if len(t.Link) == 0 {
		http.Error(w, "El link de referencia es requqerido", 400)
		return
	}
	if len(t.Codigo) < 10 {
		http.Error(w, "Inserte el Codigo de verificacion Que te genera la Empresa ToqueLaik de 10 digitos", 400)
		return
	}

	documento, encontrado, _ := bd.ChequeoYaNombreVideo(t.Nombre)
	if encontrado == true {
		http.Error(w, "Ya existe un video con ese Nombre", 400)
		return
	}
	_, status, err := bd.InsertarRegistroVideo(t)
	if err != nil {
		http.Error(w, "Ocurrio un Error al Resgistrar el Video"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el Video ", 400)
		return
	}

	jwtkey, err := jwt.GeneroJWTVideo(documento)
	if err != nil {
		http.Error(w, "Ocurrio un Error al intentar generar el Token correspndiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
