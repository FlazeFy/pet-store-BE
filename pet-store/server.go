package main

import (
	"pet-store/packages/database"
	"pet-store/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))
}
