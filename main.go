package main

import "log"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	InitRootDirectory("/var/tmp/alex_config")

	StartHttpServer("0.0.0.0:8080")
}
