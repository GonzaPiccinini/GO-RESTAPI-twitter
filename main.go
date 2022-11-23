package main

import (
	"log"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("SIN CONEXION A LA BASE DE DATOS")
		return
	}

	handlers.Handlers()
}