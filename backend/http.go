package backend

import (
	"net/http"
)

func ServeMux(staticFilePath string) *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir(staticFilePath)))
	serveMux.Handle("/rpc/json", RpcServer())
	return serveMux
}
