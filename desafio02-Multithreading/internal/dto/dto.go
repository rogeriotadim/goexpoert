package dto

import "errors"

type ApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func NewApiCep(retorno interface{}) (apiCep ApiCep, err error) {
	apiCep, ok := retorno.(ApiCep)
	if ok {
		return
	}
	return apiCep, errors.New("falha na conversção para APICEP")
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCEP(retorno interface{}) (viaCep ViaCep, err error) {
	viaCep, ok := retorno.(ViaCep)
	if ok {
		return
	}
	return viaCep, errors.New("falha na conversção para APICEP")
}
