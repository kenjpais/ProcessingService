package main

import (
	"imgservice/internal/api"
	"log"
)

func main(){
	log.Print("Starting server")

	api.StartServer()
}