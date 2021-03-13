package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pablo6911/toquemarca/bd"
	"github.com/pablo6911/toquemarca/models"
)

//GraboVideo permite grabar el mensaje en la BD
func GraboVideo(w http.ResponseWriter, r *http.Request) {
	var mensaje models.VideoMensaje
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboVideo{
		UserID:        IDUsuario,
		NombreVideo:   mensaje.NombreVideo,
		Fecha:         time.Now(),
		Email:         mensaje.Email,
		Codigo:        mensaje.Codigo,
		Link:          mensaje.Link,
		NombreEmpresa: mensaje.NombreEmpresa,
	}
	_, status, err := bd.InsertoVideo(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro reintente nuevamente"+err.Error(), 400)
	}
	if status == false {
		http.Error(w, "No se a logrado insertar los datos", 400)

	}
	w.WriteHeader(http.StatusCreated)
}
