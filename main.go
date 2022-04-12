package main

import (
	"log"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
	}

	handlers.Manejadores()

}
