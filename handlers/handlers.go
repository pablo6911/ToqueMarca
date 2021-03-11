package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pablo6911/toquemarca/middlew"
	"github.com/pablo6911/toquemarca/routers"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto, el Handler y pongo a escuchar al servidor---------
func Manejadores() {
	router := mux.NewRouter()
	//Registro video y marca
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/registroVideo", middlew.ChequeoBD(routers.Video)).Methods("POST")

	//login marca
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	//ver perfil video y marac
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.Verperfil))).Methods("GET")

	//modificar perfil video y maraca
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificoPerfil))).Methods("Put")

	//Subir foto marca
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(routers.SubirAvatar)).Methods("POST")

	//subir video -->vieo

	//obtener avatar marca
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
