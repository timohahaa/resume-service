package main

import (
	"fmt"
	"log"

	"github.com/timohahaa/resume-service/config"
)

func main() {
	config, err := config.NewConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", config)
}
