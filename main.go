package main

import "log"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	InitFiles("/home/sasha/alexconfig")

	StartHttpServer("0.0.0.0:8080")
}
