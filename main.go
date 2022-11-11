package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/megalania/erebus/pkg/server"
)

var (
	host      string
	port      int
	version   string
	timestamp string
)

func init() {
	flag.StringVar(
		&host,
		"host",
		"0.0.0.0",
		"Specify the serving host",
	)
	flag.IntVar(
		&port,
		"port",
		8080,
		"Specify the serving port",
	)
	flag.Parse()
}

func main() {
	log.Printf(
		"Erebus version %v (%v)",
		version,
		timestamp,
	)
	addr := fmt.Sprintf(
		"%v:%v",
		host,
		port,
	)
	log.Printf(
		"Starting server at %v",
		addr,
	)
	log.Println(
		"Stop the server with CONTROL-C",
	)
	err := server.Run(
		addr,
	)
	if err != nil {
		log.Fatal(err)
	}
}
