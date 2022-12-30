package adapter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/dto"
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/infra"
)

const expected = "02206-000"
const expectedTimeout = "context deadline exceeded"
const urlApiCep = "https://cdn.apicep.com/file/apicep/" + expected + ".json"
const urlViaCep = "http://viacep.com.br/ws/" + expected + "/json"

func TestApiCep(t *testing.T) {
	adapter := NewAdapter(infra.CreateContext(time.Second * 30))
	retorno, err := adapter.GetCep(urlApiCep)
	assert.Nil(t, err)
	cep, err := dto.NewApiCep(retorno)
	assert.Nil(t, err)
	assert.Equal(t, expected, cep.Code)
}

func TestApiCepTimeout(t *testing.T) {
	adapter := NewAdapter(infra.CreateContext(time.Nanosecond))
	_, err := adapter.GetCep(urlApiCep)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expectedTimeout)
}

func TestViaCep(t *testing.T) {
	adapter := NewAdapter(infra.CreateContext(time.Second * 30))
	retorno, err := adapter.GetCep(urlViaCep)
	assert.Nil(t, err)
	cep, err := dto.NewViaCEP(retorno)
	assert.Nil(t, err)
	assert.Equal(t, expected, cep.Cep)
}

func TestViaCepTimeout(t *testing.T) {
	adapter := NewAdapter(infra.CreateContext(time.Nanosecond))
	_, err := adapter.GetCep(urlViaCep)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expectedTimeout)
}
