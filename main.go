package main

import (
	"flag"
	"fmt"
	"github.com/cwahbong/jg/backend"
	"log"
	"net/http"
)

const (
	defaultPort           = 80
	defaultStaticFilePath = "./static/app/"
)

type Args struct {
	Port           uint
	StaticFilePath string
}

func parseArgs() *Args {
	var args Args
	flag.UintVar(&args.Port, "p", defaultPort, "Specify the port.")
	flag.StringVar(&args.StaticFilePath, "s", defaultStaticFilePath, "Specify the static file path.")
	flag.Parse()
	return &args
}

func main() {
	args := parseArgs()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", args.Port),
		Handler: backend.ServeMux(args.StaticFilePath),
	}
	log.Fatal(server.ListenAndServe())
}
