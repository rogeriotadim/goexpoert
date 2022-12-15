package main

import (
	"testing"
)

func TestGetCotacaoFromWeb(t *testing.T) {
	cotacao, err := GetCotacao()
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
	if cotacao.CodeIn != CODEIN {
		t.Errorf("Expected %s but got %s", CODEIN, cotacao.CodeIn)
	}
}
