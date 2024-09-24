package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("sleep")
		time.Sleep(5 * time.Second)
	}
}
