# DataStar lsp in go

build the binary

```
go build .
```

```
-- in nvim
vim.api.nvim_create_autocmd("FileType", {
	pattern = { "html" },
	callback = function(args)
		vim.lsp.start({
			name = "datastar lsp",
			cmd = { "<path/to/your/binary>" },
			root_dir = vim.fs.root(args.buf, { ".git", "go.mod" }) or vim.fn.getcwd(),
		})
	end,
})
```

[Data Star](https://data-star.dev/)

## Project Structure

```
.
├── go.mod
├── go.sum
├── handler
│   ├── completion.go
│   ├── datastar
│   │   ├── attributes
│   │   │   ├── data-bind.md
│   │   │   ├── data-on.md
│   │   │   └── data-*.md
│   │   └── data-on
│   │       ├── click.md
│   │       └── *.md
│   ├── handler.go
│   └── hover.go
├── LICENSE
├── main.go
└── README.md
```
