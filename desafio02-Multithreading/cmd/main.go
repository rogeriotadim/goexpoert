package main

import (
	"fmt"
	"time"

	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/dto"
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/infra"
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/infra/adapter"
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/usecase"
)

func main() {
	cep := "01040-000"
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	urlApi := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	urlVia := "http://viacep.com.br/ws/" + cep + "/json"
	adapterApi := adapter.NewAdapter(infra.CreateContext(time.Second * 1))
	adapterVia := adapter.NewAdapter(infra.CreateContext(time.Second * 1))
	ucApi := usecase.NewGetCepUsecase(adapterApi, urlApi)
	ucVia := usecase.NewGetCepUsecase(adapterVia, urlVia)

	go getCep(ucApi, c1)
	go getCep(ucVia, c2)

	select {
	case msg := <-c1:
		cep, err := dto.NewApiCep(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from APICEP: %v\n", cep)
	case msg := <-c2:
		cep, err := dto.NewViaCEP(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from VIACEP: %v\n", cep)
	case <-time.After(time.Second * 2):
		println("timeout")
	}
}

func getCep(uc usecase.GetCepUsecase, ch chan interface{}) (err error) {
	retorno, err := uc.GetCep()
	if err != nil {
		return
	}
	ch <- retorno
	return
}
