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

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	//router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	//router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.Verperfil))).Methods("GET")
	//router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificoPerfil))).Methods("Put")

	//router.HandleFunc("/subirAvatar", middlew.ChequeoBD(routers.SubirAvatar)).Methods("POST")
	//router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
