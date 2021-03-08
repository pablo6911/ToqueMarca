package main

import (
	"log"

	"github.com/pablo6911/toquemarca/bd"
	"github.com/pablo6911/toquemarca/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
