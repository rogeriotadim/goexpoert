package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rogeriotadim/goexpoert/desafio01-client-server-api/server"
)

const PORT = "8080"

func main() {
	http.HandleFunc("/cotacao", HandlerGetCotacao)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func HandlerGetCotacao(w http.ResponseWriter, r *http.Request) {
	log.Println("####### - HandlerGetCotacao")
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, time.Millisecond*200)
	ctxDB, cancelDB := context.WithTimeout(ctxParent, time.Millisecond*100)
	cotacao, err := server.GetCotacao(ctx, cancel, ctxDB, cancelDB)
	if err != nil {
		log.Printf("####### - HandlerGetCotacao - error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}
