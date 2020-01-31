package main

import (
	"htz/sutra/admin-server/config"
	_ "htz/sutra/admin-server/rest/controller"
	"htz/sutra/common/database"
	"htz/sutra/common/server"
)

func main() {
	database.DefaultDB.Start(config.DefaultConfig.MongoURI)
	server.DefaultServer.Start(config.DefaultConfig.HTTPAddress)
}
