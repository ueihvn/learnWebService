package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

type UUID struct {
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)

	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Fprintf(w, fmt.Sprintf("%x", b))

}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
