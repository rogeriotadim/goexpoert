package client

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCotacao(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, time.Millisecond*3000)
	cotacao, err := GetCotacao(ctx, cancel)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if cotacao.Valor == "" {
		t.Errorf("Expecting a value but get \"\"")
	}
}
func TestCotacaoTimeout(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, time.Millisecond*3)
	_, err := GetCotacao(ctx, cancel)
	expected := "context deadline exceeded"
	if !strings.Contains(fmt.Sprintf("err: %v\n", err), expected) {
		t.Errorf("expected \"%s\"", expected)
	}
}
