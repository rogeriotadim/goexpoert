package usecase

import (
	"testing"
	"time"

	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/dto"
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/infra"
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/infra/adapter"
	"github.com/stretchr/testify/assert"
)

func TestGetCep(t *testing.T) {
	expected := "02206-000"
	adapter := adapter.NewAdapter(infra.CreateContext(time.Second * 15))
	retorno, err := adapter.GetCep("http://viacep.com.br/ws/" + expected + "/json")
	assert.Nil(t, err)
	cep, err := dto.NewViaCEP(retorno)
	assert.Nil(t, err)
	assert.Equal(t, expected, cep.Cep)
}
