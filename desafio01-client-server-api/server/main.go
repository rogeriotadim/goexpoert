package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

const (
	URL_ECONOMIA="https://economia.awesomeapi.com.br/json/last/"
	CODE="USD"
	CODEIN="BRL"
	EP="cotacao"
	PORT="8080"
)

type Symbol struct {
	Symbol Cotacao `json:"USDBRL"`
}
type Cotacao struct {
	Id string  `json:"id"`
	Code string `json:"code"`
	CodeIn string `json:"codein"`
	Name string `json:"name"`
	High string `json:"high"`
	Low string `json:"low"`
	VarBid string `json:"varBid"`
	PctChange string `json:"pctChange"`
	Bid string `json:"bid"`
	Ask string `json:"ask"`
	Timestamp string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func NewCotacao(code string, codein string, name string, high string, low string, varBid string, pctChange string, bid string, ask string, timestamp string, create_date string) (cotacao Cotacao) {
	cotacao = Cotacao {
		Code: code,
		CodeIn: codein,
		Name: name,
		High: high,
		Low: low,
		VarBid: varBid,
		PctChange: pctChange,
		Bid: bid,
		Ask: ask,
		Timestamp: timestamp,
		CreateDate: create_date,
	}
	return cotacao
}

func GetCotacao() (cotacao Cotacao, err error){
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
	var data Symbol
	err = json.Unmarshal(response, &data)
	if err != nil {
		return cotacao, fmt.Errorf("parse response error: %v", err)
	}
	data.Symbol.Id = uuid.New().String()
	cotacao = data.Symbol
	return
}

func main()  {
	cotacao, err := GetCotacao()
	if err != nil {
		panic(err)
	}
	fmt.Println(cotacao)
}