package handler

import (
	"context"
	"strings"

	"github.com/owenrumney/go-lsp/lsp"
	"github.com/sumit-poudel/datastar-lsp/util"
)

func (h *Handler) Hover(_ context.Context, params *lsp.HoverParams) (*lsp.Hover, error) {
	document, ok := h.documents.Get(params.TextDocument.URI)
	if !ok {
		return nil, nil
	}
	line, ok := document.Line(int(params.Position.Line))
	word := util.WordAt(line, int(params.Position.Character))
	if word == "" {
		return nil, nil
	}

	// action is found
	attr, action, found := strings.Cut(word, ":")
	if found {
		attribute, ok := Attributes[attr].Options[action]
		if !ok {
			return nil, nil
		}
		return &lsp.Hover{
			Contents: lsp.NewHoverContents(lsp.Markdown, attribute.Desc),
		}, nil
	}

	// attribute is found
	attribute, ok := Attributes[word]
	if ok {
		return &lsp.Hover{
			Contents: lsp.NewHoverContents(lsp.Markdown, attribute.Desc),
		}, nil
	}

	// if non of the abover is found
	return nil, nil
}
