package main

import (
	"freepass-bcc/infrastucture"
	"freepass-bcc/infrastucture/database"
)

func main() {
	infrastucture.LoadEnv()
	database.ConnectToDB()
}
