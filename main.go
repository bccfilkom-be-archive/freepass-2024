package main

import (
	config "github.com/AkbarFikri/freepassBCC-2024/configs"
	"github.com/AkbarFikri/freepassBCC-2024/database"
	routers "github.com/AkbarFikri/freepassBCC-2024/routes"
)

func main() {
	config.SetupConfig()
	database.Database()
	database.Migrate()

	router := routers.SetupRoute()
	router.Run()
}
