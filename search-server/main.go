package main

import (
	"htz/sutra/common/server"
	"htz/sutra/search-server/config"
	_ "htz/sutra/search-server/rest/controller"
)

func main() {
	server.DefaultServer.Start(config.DefaultConfig.HTTPAddress)
}
