package main

import "testing"

func TestSaveCotacao(t *testing.T) {
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
		"2022-12-15 14:51:33",
	)
	err := SaveCotacao(cotacao)
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
}