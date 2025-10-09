package server

import (
	"log"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	var s Server

	http.Handle("/hi", s)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
