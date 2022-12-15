package main

import (
	"testing"
)

func TestNewCotacao(t *testing.T) {
	expectedCreateDate := "2022-12-15 14:51:00"
	cotacao := NewCotacao(
		"USD",
		"BRL",
		"Dólar Americano/Real Brasileiro",
		"5.3471",
		"5.2792",
		"0.0446",
		"0.84",
		"5.322",
		"5.3245",
		"1671124497",
		expectedCreateDate,
	)
	if cotacao.CreateDate != expectedCreateDate {
		t.Errorf("Expected %s but got %s", expectedCreateDate, cotacao.CreateDate)
	}
}

func TestGetCotacaoFromWeb(t *testing.T) {
	cotacao, err := GetCotacao()
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
	if cotacao.CodeIn != CODEIN {
		t.Errorf("Expected %s but got %s", CODEIN, cotacao.CodeIn)
	}
}