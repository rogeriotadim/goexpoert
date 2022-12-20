package client

import (
	"fmt"
	"os"
)

func SaveCotacao(cotacao CotacaoDtoIn)  (err error){
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file error: %v", err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar: {%s}", cotacao.Valor))
	return

}