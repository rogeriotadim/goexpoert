package server

import (
	"context"
	"testing"
	"time"
)

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
	ctxParent := context.Background()
	ctxDB, cancelDB := context.WithTimeout(ctxParent, time.Millisecond * 10)
	defer cancelDB()

	err := SaveCotacao(ctxDB, cotacao)
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
}