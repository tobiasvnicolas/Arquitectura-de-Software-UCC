package main

import (
	"Arquitectura-de-Software-UCC/backend/app"
	"Arquitectura-de-Software-UCC/backend/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()

}
