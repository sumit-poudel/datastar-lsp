# Contributing to datastar-lsp

Thanks for your interest in contributing! `datastar-lsp` is a Language Server
Protocol implementation for [Datastar](https://data-star.dev), written in Go.
This document covers how the project is organized and how to make changes.

## Current Priority

Right now the biggest help is filling in missing attribute docs under handler/datastar/. assets/data-attributes.json has the full spec (every attribute + description + doc link) — just not every one has a markdown file yet.

To add one: find an attribute in the JSON that's missing a file, write a .md for it using the JSON's description/link as a starting point (rewording is fine, just keep it accurate), and open a PR. One attribute per PR is easiest to review. 

## Prerequisites

- Go (see `go.mod` for the minimum version)
- An editor with LSP support (Neovim, VS Code, etc.) for manual testing

## Project layout

```
.
├── assets/                      # Reference/source data
│   ├── data-attributes.json     # Legacy/source attribute spec
│   ├── datastar.injection.tmLanguage.json
│   └── language-data.json
├── handler/                      # LSP request handlers
│   ├── datastar/                 # Datastar-specific docs & completions
│   │   ├── attributes/           # One markdown file per attribute
│   │   │   ├── data-bind.md
│   │   │   └── data-on.md
│   │   └── data-on/              # Nested docs for attribute variants
│   │       └── click.md          # e.g. data-on:click
│   ├── completion.go              # textDocument/completion
│   ├── datastar.go                # Datastar spec loading/lookup
│   ├── handler.go                 # Handler struct, Initialize, Shutdown
│   └── hover.go                   # textDocument/hover
├── util/                          # Shared helpers (position/offset math, etc.)
└── main.go                        # Server entrypoint (go-lsp wiring)
```

## Adding or updating a Datastar attribute

Attribute documentation lives as markdown files under
`handler/datastar/attributes/`, one file per attribute. Attributes with
named variants (like `data-on:click`, `data-on:keydown`) are nested under a
directory named after the base attribute, with one file per variant (e.g.
`handler/datastar/data-on/click.md`).

Each file is plain markdown, no frontmatter — but lets keep this clean shape 

```markdown
data-on:click

Attaches an event listener to an element, executing an expression whenever the event is triggered.

Notes

    <button data-on:click="$foo = ''">Reset</button>

[DataStar Refrence](https://data-star.dev/reference/attributes#data-on)
```

When you add or edit one of these files:

1. **First line** is the attribute name as it appears in markup (e.g.
   `data-bind`, `data-on:click`). No heading syntax (`#`) — just the bare
   name on its own line.
2. **Description paragraph** follows after a blank line. Keep it concise —
   it's shown directly in hover popups and completion documentation panels
   in the user's editor.
3. **Optional `Notes` section** — for attributes where a short usage example
   helps, add a `Notes` line followed by an indented code block.
4. **Closing reference link** — always end the file with a markdown link
   back to the relevant page on the official
   [Datastar docs](https://data-star.dev/reference/attributes) site.
5. These files are loaded via `go:embed` at build time — no separate build
   step is required, just rebuild the binary.

## Building

```sh
go build -o datastar-lsp .
```

## Testing locally with Neovim

Point `vim.lsp.start` at your locally built binary, e.g.:

```lua
vim.api.nvim_create_autocmd("FileType", {
  pattern = { "html" }, -- target you file not me gang 
  callback = function(args)
    vim.lsp.start({
      name = "datastar-lsp",
      cmd = { "/path/to/your/datastar-lsp" },
      root_dir = vim.fs.root(args.buf, { ".git", "go.mod" }) or vim.fn.getcwd(),
    })
  end,
})
```

Check `:checkhealth vim.lsp` | `:Lspinfo` to confirm the server attached, and watch for conflicts if
other language servers (e.g. `htmx-lsp`) are also enabled for the same
filetype — LSP hover/completion results are merged across all attached
clients, so an unrelated server can interfere with what you see while
testing.

## Running tests

I havent actually dont written any test cases for now

## Debug session

open browser on localhost:7100 for debug session

```
	srv := server.NewServer(h)

    // add this instead
	srv := server.NewServer(h, server.WithDebugUI(":7100"))
```

[Click here for more info](https://github.com/owenrumney/go-lsp.git)

## Code style

- Run `gofmt`/`go vet` before opening a PR.
- Keep handler methods small and focused — one LSP method per file where
  reasonable (see `completion.go`, `hover.go`).
- Prefer returning empty results (`&lsp.CompletionList{Items: []lsp.CompletionItem{}}`)
  over `nil` where the LSP spec expects a list, to avoid null-related issues
  on the client side.
- Use util/ for all utils and others for the the corresponding named files

## Submitting a change

1. Fork the repo and create a branch for your change.
2. Make your changes, including tests where applicable.
3. Ensure `go build ./...` and `go test ./...` both pass.
4. Open a pull request with a clear description of what changed and why.
5. Releases are automated via GitHub Actions (`.github/workflows/release.yml`)
   and GoReleaser (`.goreleaser.yaml`) — you don't need to worry about
   versioning or publishing as a contributor.

## Questions

If anything here is unclear, feel free to open an issue — this project is
still evolving and this document will grow along with it.
