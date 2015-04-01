package main

import "log"
import "encoding/json"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	InitRootDirectory("/home/sasha/alexconfig")

	var value interface{}

	value, ok, err := GetNode("fff")
	if !ok {
		log.Panic(err)
	}

	txt, er := json.Marshal(value)
	if er != nil {
		log.Panic(er)
	}

	log.Println(string(txt))

	StartHttpServer("0.0.0.0:8080")
}
