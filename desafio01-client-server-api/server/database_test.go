package server

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestSaveCotacao(t *testing.T) {
	cotacao := CotacaoTest
	CotacaoTest.Id = uuid.New().String()
	ctxParent := context.Background()
	ctxDB, cancelDB := context.WithTimeout(ctxParent, time.Millisecond*1000)
	defer cancelDB()

	err := SaveCotacao(ctxDB, cotacao)
	if err != nil {
		t.Fatalf("Erro: %v", err)
	}
}

func TestSaveCotacaoTimeout(t *testing.T) {
	cotacao := CotacaoTest
	CotacaoTest.Id = uuid.New().String()
	ctxParent := context.Background()
	ctxDB, cancelDB := context.WithTimeout(ctxParent, time.Nanosecond*1)
	defer cancelDB()

	err := SaveCotacao(ctxDB, cotacao)
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("Erro: %v", err)
	}
}
