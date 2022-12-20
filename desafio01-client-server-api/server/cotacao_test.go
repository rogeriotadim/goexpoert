package server

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
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
	if cotacao.Id == "" {
		t.Error("Expected Id not null")
	}
}

func TestGetCotacaoFromWeb(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, time.Millisecond*2000)
	ctxDB, cancelDB := context.WithTimeout(ctxParent, time.Millisecond*1000)
	cotacao, err := GetCotacao(ctx, cancel, ctxDB, cancelDB)
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
	if cotacao.Valor == "" {
		t.Error("Expected a value but got empty")
	}
}

func TestGetCotacaoFromWebTimeoutHttp(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, time.Millisecond*20)
	ctxDB, cancelDB := context.WithTimeout(ctxParent, time.Millisecond*1000)
	_, err := GetCotacao(ctx, cancel, ctxDB, cancelDB)
	expected := "context deadline exceeded"
	if !strings.Contains(fmt.Sprintf("err: %v\n", err), expected) {
		t.Errorf("expected \"%s\"", expected)
	}
}

func TestGetCotacaoFromWebTimeoutDB(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, time.Millisecond*2000)
	ctxDB, cancelDB := context.WithTimeout(ctxParent, time.Millisecond*1)
	_, err := GetCotacao(ctx, cancel, ctxDB, cancelDB)
	if err == nil {
		t.Fatal("error should not be nil")
	}
	expected := "symbol persisting cancelled"
	errString := fmt.Sprintf("err: %v\n", err)
	if !strings.Contains(errString, expected) {
		t.Errorf("expected \"%s\" but \"%s\"", expected, errString)
	}
}
