package main

import (
	"testing"
)

func TestGetCotacaoFromWeb(t *testing.T) {
	cotacao, err := GetCotacao()
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
	if cotacao.Valor == "" {
		t.Error("Expected a value but got empty")
	}
}
