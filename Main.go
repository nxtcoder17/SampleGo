package main

import (
	"SampleCrud/db"
	"SampleCrud/routes"
)

func main() {
	db.Init()
	routes.Init()
}
