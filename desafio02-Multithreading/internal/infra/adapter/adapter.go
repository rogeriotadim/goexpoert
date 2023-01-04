package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/dto"
)

type Adapter struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewAdapter(ctx context.Context, cancel context.CancelFunc) (adapter Adapter) {
	return Adapter{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *Adapter) GetCep(url string) (retorno interface{}, err error) {
	req, err := http.NewRequestWithContext(a.ctx, "GET", url, nil)
	if err != nil {
		return
	}
	request, err := http.DefaultClient.Do(req)
	if err != nil {
		return retorno, fmt.Errorf("request error: %v", err)
	}
	if request.StatusCode > 399 {
		return retorno, fmt.Errorf("request error: %v", err)
	}
	defer request.Body.Close()
	response, err := io.ReadAll(request.Body)
	if err != nil {
		return retorno, fmt.Errorf("get response error: %v", err)
	}
	cep, err := returnInterface(response)
	if err != nil {
		return retorno, err
	}
	return cep, nil
}

func returnInterface(response []byte) (i interface{}, err error) {
	respString := string(response)
	if strings.Contains(respString, "statusText") {
		var cep dto.ApiCep
		err = json.Unmarshal(response, &cep)
		if err != nil {
			return
		}
		return cep, nil
	}
	if strings.Contains(respString, "logradouro") {
		var cep dto.ViaCep
		err = json.Unmarshal(response, &cep)
		if err != nil {
			return
		}
		return cep, nil
	}
	return i, fmt.Errorf("falha na conversa da interface para DTO")
}
