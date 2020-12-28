package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ohshane71/go-http-tut/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)

	http.ListenAndServe(":9090", sm)
}
