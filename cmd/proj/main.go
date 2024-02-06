package main

import (
	"fmt"
	"log"

	"github.com/Ozzy-ZY/proj/internal/helloworld"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)
	message , err := helloworld.Hello("")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)
}
