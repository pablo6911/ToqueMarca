package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pablo6911/toquemarca/bd"
	"github.com/pablo6911/toquemarca/models"
)

//Video es la func para crear en la BD el registro de user-------
func Video(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "Codigo de verificacion Que te genera la Empresa ToqueLaik", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaNombreVideo(t.Nombre)
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
	w.WriteHeader(http.StatusCreated)
}
