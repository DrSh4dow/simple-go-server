package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("start %s\n", time.Now().String())
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello\n")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	_ = http.ListenAndServe(":8000", &l)
}
