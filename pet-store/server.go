package main

import (
	// "pet-store/factories"
	"pet-store/packages/database"
	"pet-store/routes"
	// "pet-store/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))

	// Run Seeders
	// factories.Factory()
}
