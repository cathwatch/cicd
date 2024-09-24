package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("sleep two")
		time.Sleep(5 * time.Second)
	}
}
