package config

import (
	"log"
	"os"
	"strconv"
)

var (
	PORT = 0
)

func Load() {
	var err error
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		PORT = 3000
	}

}
