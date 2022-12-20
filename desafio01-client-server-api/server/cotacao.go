package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const (
	URL_ECONOMIA = "https://economia.awesomeapi.com.br/json/last/"
	CODE         = "USD"
	CODEIN       = "BRL"
	EP           = "cotacao"
	PORT         = "8080"
)

type CotacaoDtoOut struct {
	Valor string `json:"valor"`
}

type Symbol struct {
	Symbol Cotacao `json:"USDBRL"`
}
type Cotacao struct {
	Id         string `json:"id"`
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func NewCotacao(code string, codein string, name string, high string, low string, varBid string, pctChange string, bid string, ask string, timestamp string, create_date string) (cotacao Cotacao) {
	cotacao = Cotacao{
		Id:         uuid.New().String(),
		Code:       code,
		CodeIn:     codein,
		Name:       name,
		High:       high,
		Low:        low,
		VarBid:     varBid,
		PctChange:  pctChange,
		Bid:        bid,
		Ask:        ask,
		Timestamp:  timestamp,
		CreateDate: create_date,
	}
	return cotacao
}

func GetCotacao(ctx context.Context, cancel context.CancelFunc, ctxDB context.Context, cancelDB context.CancelFunc) (cotacaoOut CotacaoDtoOut, err error) {
	log.Println("####### - GetCotacao")
	defer cancel()
	defer cancelDB()

	ep := URL_ECONOMIA + CODE + "-" + CODEIN
	log.Println("####### - GetCotacao - ep: ", ep)
	req, err := http.NewRequestWithContext(ctx, "GET", ep, nil)
	if err != nil {
		return
	}
	request, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if err != nil {
		return cotacaoOut, errors.New("request failed")
	}
	defer request.Body.Close()
	response, err := io.ReadAll(request.Body)
	if err != nil {
		return cotacaoOut, fmt.Errorf("get response error: %v", err)
	}
	var data Symbol
	err = json.Unmarshal(response, &data)
	if err != nil {
		return cotacaoOut, fmt.Errorf("parse response error: %v", err)
	}
	data.Symbol.Id = uuid.New().String()
	cotacao := data.Symbol
	err = SaveCotacao(ctxDB, cotacao)
	if err != nil {
		return cotacaoOut, fmt.Errorf("persist cotacao error: %v", err)
	}
	cotacaoOut.Valor = cotacao.Bid
	return
}
