package main

import (
	_ "htz/sutra/admin-server/rest/controller"
	"htz/sutra/admin-server/rest/server"
)

func main() {
	//database.DefaultDB.Start(config.DefaultConfig.MongoURI)
	server.DefaultServer.Start()
}
