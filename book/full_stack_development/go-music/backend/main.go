package main

import (
	"gomusic/rest"
	"log"
)

func main() {
	log.Println("main log")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
