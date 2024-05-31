package main

import (
	"log"
	"net/http"
)

func (botcfg *BotApiConfig) HandleReadiness(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Fatal(err)
	}

}
