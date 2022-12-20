package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CotacaoDtoIn struct {
	Valor string `json:"valor"`
}

func GetCotacao(ctx context.Context, cancel context.CancelFunc) (cotacaoIn CotacaoDtoIn, err error) {
	defer cancel()

	ep := "http://localhost:8080/cotacao"
	req, err := http.NewRequestWithContext(ctx, "GET", ep, nil)
	if err != nil {
		return
	}
	request, err := http.DefaultClient.Do(req)
	if err != nil {
		return cotacaoIn, fmt.Errorf("request error: %v", err)
	}
	if request.StatusCode > 399 {
		return cotacaoIn, fmt.Errorf("request error: %v", err)
	}
	defer request.Body.Close()
	response, err := io.ReadAll(request.Body)
	if err != nil {
		return cotacaoIn, fmt.Errorf("get response error: %v", err)
	}
	err = json.Unmarshal(response, &cotacaoIn)
	if err != nil {
		return cotacaoIn, fmt.Errorf("parse response error: %v", err)
	}
	return
}
