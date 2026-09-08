package handler

import (
	"context"

	"github.com/owenrumney/go-lsp/document"
	"github.com/owenrumney/go-lsp/lsp"
	"github.com/owenrumney/go-lsp/server"
)

type Handler struct {
	documents *document.Store
	client    *server.Client
}

func New() *Handler {
	return &Handler{
		documents: document.NewStore(),
	}
}

func (h *Handler) Initialize(_ context.Context, _ *lsp.InitializeParams) (*lsp.InitializeResult, error) {
	return &lsp.InitializeResult{
		ServerInfo: &lsp.ServerInfo{Name: "datastar lsp", Version: "0.1.0"},
	}, nil
}

func (h *Handler) Shutdown(_ context.Context) error { return nil }

func (h *Handler) DidOpen(_ context.Context, params *lsp.DidOpenTextDocumentParams) error {
	_, err := h.documents.Open(params)
	return err
}

func (h *Handler) DidChange(_ context.Context, params *lsp.DidChangeTextDocumentParams) error {
	_, err := h.documents.Change(params)
	return err
}

func (h *Handler) DidClose(_ context.Context, params *lsp.DidCloseTextDocumentParams) error {
	h.documents.Close(params)
	return nil
}
