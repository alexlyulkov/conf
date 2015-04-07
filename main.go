package main

import "log"
import "github.com/alexlyulkov/conf/conf"
import "github.com/alexlyulkov/conf/server"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	conf.InitRootDirectory("/var/tmp/alex_config")

	server.StartHttpServer("0.0.0.0:8080")
}
