package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ohshane71/go-http-tut/data"
)

// do not name productsHandler
// it's already in the package name
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
