package main

import (
	"log"
	"os"
	"time"
)

const shutdownTimeout = 10 * time.Second

var (
	buildTimestamp string
	buildVersion   string
)

func init() {
	tiz := os.Getenv(
		"TZ",
	)
	loc, err := time.LoadLocation(
		tiz,
	)
	if err != nil {
		log.Fatalln(
			err.Error(),
		)
	}
	time.Local = loc
}

func main() {
	log.Println(
		"Hello, world!",
	)
}
