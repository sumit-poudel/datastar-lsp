package main

import (
	"context"
	"log"

	"github.com/owenrumney/go-lsp/server"
	"github.com/sumit-poudel/datastar-lsp/handler"
)

func main() {
	h := handler.New()
	srv := server.NewServer(h, server.WithDebugUI(":7100"))
	if err := srv.Run(context.Background(), server.RunStdio()); err != nil {
		log.Fatal(err)
	}
}
