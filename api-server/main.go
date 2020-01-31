package main

import (
	_ "htz/sutra/search-server/search"
	"log"
	"htz/sutra/common/database"
	"htz/sutra/api-server/config"
	"htz/sutra/common/server"
	_ "htz/sutra/api-server/rest/controller"
)

func init()  {
	log.SetFlags(log.Lshortfile)
}

func main() {
	database.DefaultDB.Start(config.DefaultConfig.MongoURI)
	server.DefaultServer.Start(config.DefaultConfig.HTTPAddress)
}
