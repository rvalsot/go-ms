package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	loog *log.Logger
}

func NewGoodbye(loog *log.Logger) *Goodbye {
	return &Goodbye{loog}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Seee ya! \n"))
}
