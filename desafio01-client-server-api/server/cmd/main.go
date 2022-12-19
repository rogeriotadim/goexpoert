package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rogeriotadim/goexpoert/desafio01-client-server-api/server"
)

const PORT="8080"

func main()  {
	http.HandleFunc("/cotacao", HandlerGetCotacao)
	log.Fatal(http.ListenAndServe(":" + PORT, nil))
}

func HandlerGetCotacao(w http.ResponseWriter, r *http.Request){
	cotacao, err := server.GetCotacao()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}