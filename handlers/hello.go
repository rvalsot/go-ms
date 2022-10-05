package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	loog *log.Logger
}

func NewHellow(loog *log.Logger) *Hello {
	return &Hello{loog}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.loog.Println("Hellow wourld!")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "Ooopsy, a request error ocurred", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello, %s \n", data)

}
