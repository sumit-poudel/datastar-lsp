package handler

import (
	"context"
	"fmt"

	"github.com/owenrumney/go-lsp/lsp"
)

func (h *Handler) Completion(_ context.Context, _ *lsp.CompletionParams) (*lsp.CompletionList, error) {
	kind := lsp.CompletionItemKindKeyword

	items := make([]lsp.CompletionItem, 0, len(Attributes))
	for _, attribute := range Attributes {
		items = append(items, lsp.CompletionItem{
			Label:  attribute.Name,
			Kind:   &kind,
			Detail: attribute.Desc,
		})
	}

	return &lsp.CompletionList{Items: items}, nil
}

func (h *Handler) ResolveCompletionItem(_ context.Context, item *lsp.CompletionItem) (*lsp.CompletionItem, error) {
	item.Documentation = &lsp.MarkupContent{
		Kind:  lsp.Markdown,
		Value: fmt.Sprintf("The **%s** keyword in Go.", item.Label),
	}
	return item, nil
}
