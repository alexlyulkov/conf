package main

import "log"
import "github.com/alexlyulkov/config/conf"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	conf.InitRootDirectory("/var/tmp/alex_config")

	conf.StartHttpServer("0.0.0.0:8080")
}
