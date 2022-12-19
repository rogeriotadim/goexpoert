package server

import "github.com/google/uuid"

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
		Id: uuid.New().String(),
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
