package main

import (
	"github.com/cwahbong/jg/backend"
	"log"
	"net/http"
)

const (
	DbName = "jg_test"
)

func main() {
	// TODO set static file path by parsing args
	staticFilePath := "./static/app/"

	http.Handle("/", http.FileServer(http.Dir(staticFilePath)))
	http.Handle("/rpc/json", backend.RpcServer())

	// userprofile
	server := http.Server{
		Addr: ":9765",
	}
	log.Fatal(server.ListenAndServe())
}
