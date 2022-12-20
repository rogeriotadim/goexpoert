package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rogeriotadim/goexpoert/desafio01-client-server-api/client"
)

func main() {
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, time.Millisecond*300)
	cotacao, err := client.GetCotacao(ctx, cancel)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		os.Exit(1)
	}
	err = client.SaveCotacao(cotacao)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		os.Exit(1)
	}

}
