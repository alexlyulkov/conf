package main

import (
	"flag"
	"github.com/alexlyulkov/conf/conf"
	"github.com/alexlyulkov/conf/server"
	"log"
)

type Params struct {
	ServerAddress    string
	WorkingDirectory string
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	params := getComandLineFlags()
	conf.InitRootDirectory(params.WorkingDirectory)

	server.StartHttpServer(params.ServerAddress)
}

func getComandLineFlags() *Params {
	flags := new(Params)
	flag.StringVar(&flags.ServerAddress, "address", "", "Server address (host:port)")
	flag.StringVar(&flags.WorkingDirectory, "workdir", "", "Working directory")
	flag.Parse()
	return flags
}
