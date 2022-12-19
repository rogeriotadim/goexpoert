package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/rogeriotadim/goexpoert/desafio01-client-server-api/server"
)

const (
	URL_ECONOMIA="https://economia.awesomeapi.com.br/json/last/"
	CODE="USD"
	CODEIN="BRL"
	EP="cotacao"
	PORT="8080"
)


func GetCotacao() (cotacao server.Cotacao, err error){
	ep := URL_ECONOMIA + CODE + "-" + CODEIN
	request, err := http.Get(ep)
	if err != nil {
		return cotacao, errors.New("request failed") 
	}
	defer request.Body.Close()
	response, err := io.ReadAll(request.Body)
	if err != nil {
		return cotacao, fmt.Errorf("get response error: %v", err)
	}
	var data server.Symbol
	err = json.Unmarshal(response, &data)
	if err != nil {
		return cotacao, fmt.Errorf("parse response error: %v", err)
	}
	data.Symbol.Id = uuid.New().String()
	cotacao = data.Symbol
	err = server.SaveCotacao(cotacao)
	if err != nil {
		return cotacao, fmt.Errorf("persist cotacao error: %v", err)
	}
	return
}

func main()  {
	http.HandleFunc("/", HandlerGetCotacao)
	log.Fatal(http.ListenAndServe(":" + PORT, nil))
}

func HandlerGetCotacao(w http.ResponseWriter, r *http.Request){
	cotacao, err := GetCotacao()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}