package main

import (
	"log"

	"github.com/sabe-0524/proglog/internal/server"
)

func main() {
	httpsrv := server.NewHTTPServer(":8080")
	log.Fatal(httpsrv.ListenAndServe())
}