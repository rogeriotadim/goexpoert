package usecase

import (
	"github.com/rogeriotadim/goexpoert/desafio02-Multithreading/internal/infra/adapter"
)

type GetCepUsecase struct {
	adapter adapter.Adapter
	url     string
}

func NewGetCepUsecase(adapter adapter.Adapter, url string) (usecase GetCepUsecase) {
	usecase.adapter = adapter
	usecase.url = url
	return usecase
}

func (u *GetCepUsecase) GetCep() (i interface{}, err error) {
	i, err = u.adapter.GetCep(u.url)
	if err != nil {
		return i, err
	}
	return
}
