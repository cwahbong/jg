package commands

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

type serverArgs struct {
	Port           uint
	StaticFilePath string
}

func jgServer(argstrs []string) {
	var args serverArgs
	flagSet := flag.NewFlagSet("jg-server", flag.ExitOnError)
	flagSet.UintVar(&args.Port, "p", defaultPort, "Specify the port.")
	flagSet.StringVar(&args.StaticFilePath, "s", defaultStaticFilePath, "Specify the static file path.")
	flagSet.Parse(argstrs)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", args.Port),
		Handler: backend.ServeMux(args.StaticFilePath),
	}
	log.Fatal(server.ListenAndServe())
}
