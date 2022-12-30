package adapter

import (
	"testing"
	"time"

	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/dto"
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/infra"
)

const expected = "02206-000"
const urlApiCep = "https://cdn.apicep.com/file/apicep/" + expected + ".json"
const urlViaCep = "http://viacep.com.br/ws/" + expected + "/json"

func TestApiCep(t *testing.T) {
	adapter := NewAdapter(infra.CreateContext(time.Second * 30))
	retorno, err := adapter.GetCep(urlApiCep)
	if err != nil {
		t.Error(err)
	}
	cep, err := dto.NewApiCep(retorno)
	if err != nil {
		t.Error(err)
	}
	if cep.Code != expected {
		t.Errorf("cep não confere")
	}
}

func TestViaCep(t *testing.T) {
	adapter := NewAdapter(infra.CreateContext(time.Second * 30))
	retorno, err := adapter.GetCep(urlViaCep)
	if err != nil {
		t.Error(err)
	}
	cep, err := dto.NewViaCEP(retorno)
	if err != nil {
		t.Error(err)
	}
	if cep.Cep != expected {
		t.Errorf("cep não confere")
	}
}
