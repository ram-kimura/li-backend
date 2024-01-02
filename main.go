package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"li-backend/hello"
	"net/http"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Info().Msg("info log test")

	hello.Hello()

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/header", headerHandler)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Error().Msgf("Failed to listen and serve: %v", err)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, Handler!\n"))
	if err != nil {
		log.Error().Msgf("Failed to write response: %v", err)
		return
	}
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				log.Error().Msgf("Failed to write response: %v", err)
				return
			}
		}
	}
}
